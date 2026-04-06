//go:build windows

package simconnect

import (
	"encoding/binary"
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
	"unsafe"
)

var (
	simconnectDLL = syscall.NewLazyDLL("SimConnect.dll")

	procOpen                   = simconnectDLL.NewProc("SimConnect_Open")
	procClose                  = simconnectDLL.NewProc("SimConnect_Close")
	procAddToDataDefinition    = simconnectDLL.NewProc("SimConnect_AddToDataDefinition")
	procClearDataDefinition    = simconnectDLL.NewProc("SimConnect_ClearDataDefinition")
	procRequestDataOnSimObject = simconnectDLL.NewProc("SimConnect_RequestDataOnSimObjectType")
	procSetDataOnSimObject     = simconnectDLL.NewProc("SimConnect_SetDataOnSimObject")
	procMapClientEvent         = simconnectDLL.NewProc("SimConnect_MapClientEventToSimEvent")
	procTransmitEvent          = simconnectDLL.NewProc("SimConnect_TransmitClientEvent")
	procFlightPlanLoad         = simconnectDLL.NewProc("SimConnect_FlightPlanLoad")
	procFlightSave             = simconnectDLL.NewProc("SimConnect_FlightSave")
	procRequestSystemState     = simconnectDLL.NewProc("SimConnect_RequestSystemState")
	procCameraSetRelative6DOF  = simconnectDLL.NewProc("SimConnect_CameraSetRelative6DOF")
	procAICreateParked         = simconnectDLL.NewProc("SimConnect_AICreateParkedATCAircraft")
	procAIRemoveObject         = simconnectDLL.NewProc("SimConnect_AIRemoveObject")
	procGetNextDispatch        = simconnectDLL.NewProc("SimConnect_GetNextDispatch")
)

// WinClient is the Windows SimConnect client.
type WinClient struct {
	mu        sync.Mutex
	handle    uintptr
	connected atomic.Bool
	dllPath   string

	events    map[string]uint32 // event name -> assigned ID
	nextEvent uint32
	nextDef   uint32
	nextReq   uint32

	// pending maps request IDs to channels for async response delivery.
	pending   map[uint32]chan []byte
	pendingMu sync.Mutex
}

func NewClient(dllPath string) Client {
	if dllPath != "" && dllPath != "SimConnect.dll" {
		simconnectDLL = syscall.NewLazyDLL(dllPath)
	}
	return &WinClient{
		events:    make(map[string]uint32),
		pending:   make(map[uint32]chan []byte),
		nextEvent: 1,
		nextDef:   1,
		nextReq:   1,
		dllPath:   dllPath,
	}
}

func (c *WinClient) Connect() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.connected.Load() {
		return nil
	}

	if err := simconnectDLL.Load(); err != nil {
		return fmt.Errorf("SimConnect.dll not found: %w", err)
	}

	name, _ := syscall.BytePtrFromString("InvoTalk-MCP")
	var handle uintptr

	hr, _, _ := procOpen.Call(
		uintptr(unsafe.Pointer(&handle)),
		uintptr(unsafe.Pointer(name)),
		0, 0, 0, 0,
	)
	if hr != HrOK {
		return fmt.Errorf("SimConnect_Open failed: HRESULT 0x%08X", hr)
	}

	c.handle = handle
	c.connected.Store(true)

	// Start dispatch goroutine for async responses.
	go c.dispatchLoop()

	return nil
}

func (c *WinClient) Close() {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.connected.Load() {
		procClose.Call(c.handle)
		c.connected.Store(false)
		c.handle = 0
	}
}

func (c *WinClient) IsConnected() bool {
	return c.connected.Load()
}

// dispatchLoop polls SimConnect for async responses and routes them to waiting callers.
func (c *WinClient) dispatchLoop() {
	for c.connected.Load() {
		var pData uintptr
		var cbData uint32

		hr, _, _ := procGetNextDispatch.Call(
			c.handle,
			uintptr(unsafe.Pointer(&pData)),
			uintptr(unsafe.Pointer(&cbData)),
		)
		if hr != HrOK || cbData == 0 || pData == 0 {
			time.Sleep(10 * time.Millisecond)
			continue
		}
		if cbData < 12 {
			continue
		}

		// First 4 bytes: size, next 4: version, next 4: ID (recv type)
		recvID := *(*uint32)(unsafe.Pointer(pData + 8))

		switch recvID {
		case RecvIDSimObjectDataByType:
			c.handleSimObjectData(pData, cbData)
		case RecvIDSystemState:
			c.handleSystemState(pData, cbData)
		case RecvIDQuit:
			c.connected.Store(false)
			return
		}
	}
}

// handleSimObjectData processes a SIMCONNECT_RECV_SIMOBJECT_DATA_BYTYPE message.
// Header layout: size(4) version(4) id(4) requestID(4) objectID(4) defineID(4)
// flags(4) entry(4) outof(4) defineCount(4) — data starts at offset 40.
func (c *WinClient) handleSimObjectData(pData uintptr, cbData uint32) {
	if cbData < 44 {
		return
	}
	reqID := *(*uint32)(unsafe.Pointer(pData + 12))
	dataSize := cbData - 40
	data := make([]byte, dataSize)
	for i := uint32(0); i < dataSize; i++ {
		data[i] = *(*byte)(unsafe.Pointer(pData + 40 + uintptr(i)))
	}
	c.deliver(reqID, data)
}

// handleSystemState processes a SIMCONNECT_RECV_SYSTEM_STATE message.
// String data starts at offset 24, null-terminated, up to 256 bytes.
func (c *WinClient) handleSystemState(pData uintptr, cbData uint32) {
	if cbData < 20 {
		return
	}
	reqID := *(*uint32)(unsafe.Pointer(pData + 12))
	strBytes := make([]byte, 0, 256)
	for i := uintptr(0); i < 256 && (24+i) < uintptr(cbData); i++ {
		b := *(*byte)(unsafe.Pointer(pData + 24 + i))
		if b == 0 {
			break
		}
		strBytes = append(strBytes, b)
	}
	c.deliver(reqID, strBytes)
}

// deliver routes a response payload to the waiting caller for reqID.
func (c *WinClient) deliver(reqID uint32, data []byte) {
	c.pendingMu.Lock()
	ch, ok := c.pending[reqID]
	if ok {
		delete(c.pending, reqID)
	}
	c.pendingMu.Unlock()
	if ok {
		ch <- data
	}
}

func (c *WinClient) ensureConnected() error {
	if !c.connected.Load() {
		if err := c.Connect(); err != nil {
			return fmt.Errorf("not connected to MSFS (reconnect failed: %w)", err)
		}
	}
	return nil
}

func (c *WinClient) allocReqID() uint32 {
	c.mu.Lock()
	defer c.mu.Unlock()
	id := c.nextReq
	c.nextReq++
	return id
}

func (c *WinClient) allocDefID() uint32 {
	c.mu.Lock()
	defer c.mu.Unlock()
	id := c.nextDef
	c.nextDef++
	return id
}

// waitForResponse registers a pending request and waits for the dispatch loop to deliver data.
func (c *WinClient) waitForResponse(reqID uint32, timeout time.Duration) ([]byte, error) {
	ch := make(chan []byte, 1)
	c.pendingMu.Lock()
	c.pending[reqID] = ch
	c.pendingMu.Unlock()

	select {
	case data := <-ch:
		return data, nil
	case <-time.After(timeout):
		c.pendingMu.Lock()
		delete(c.pending, reqID)
		c.pendingMu.Unlock()
		return nil, fmt.Errorf("timeout waiting for SimConnect response (request %d)", reqID)
	}
}

// isStringUnit returns true if the unit indicates a string-type SimVar.
func isStringUnit(unit string) bool {
	return unit == "string" || unit == "String"
}

// varLayout tracks the data type and byte size for each variable in a definition.
type varLayout struct {
	isString bool
	size     int // bytes: 8 for float64, 256 for string256
}

const string256Size = 256

func (c *WinClient) GetVariables(vars []VarRequest) (map[string]interface{}, error) {
	if err := c.ensureConnected(); err != nil {
		return nil, err
	}

	defID := c.allocDefID()
	reqID := c.allocReqID()

	// Define each variable and track the layout for parsing.
	layouts := make([]varLayout, len(vars))
	for i, v := range vars {
		namePtr, _ := syscall.BytePtrFromString(v.Name)
		unitPtr, _ := syscall.BytePtrFromString(v.Unit)

		dataType := uintptr(DataTypeFloat64)
		layouts[i] = varLayout{isString: false, size: 8}

		if isStringUnit(v.Unit) {
			dataType = DataTypeString256
			layouts[i] = varLayout{isString: true, size: string256Size}
			// String vars don't use units — pass NULL.
			unitPtr = nil
		}

		hr, _, _ := procAddToDataDefinition.Call(
			c.handle,
			uintptr(defID),
			uintptr(unsafe.Pointer(namePtr)),
			uintptr(unsafe.Pointer(unitPtr)),
			dataType,
			0,      // epsilon
			Unused, // datum ID: SIMCONNECT_UNUSED
		)
		if hr != HrOK {
			procClearDataDefinition.Call(c.handle, uintptr(defID))
			return nil, fmt.Errorf("AddToDataDefinition(%s) failed: 0x%08X", v.Name, hr)
		}
	}

	// Request the data. Signature: (handle, reqID, defID, radiusMeters, objectType)
	hr, _, _ := procRequestDataOnSimObject.Call(
		c.handle,
		uintptr(reqID),
		uintptr(defID),
		0,                 // dwRadiusMeters: 0 = user aircraft only
		SimObjectTypeUser, // eType: 0 = SIMCONNECT_SIMOBJECT_TYPE_USER
	)
	if hr != HrOK {
		return nil, fmt.Errorf("RequestDataOnSimObjectType failed: 0x%08X", hr)
	}

	// Wait for response from dispatch loop.
	data, err := c.waitForResponse(reqID, 5*time.Second)
	if err != nil {
		return nil, err
	}

	// Parse values from the raw bytes using the tracked layout.
	result := make(map[string]interface{}, len(vars))
	offset := 0
	for i, v := range vars {
		if offset+layouts[i].size > len(data) {
			procClearDataDefinition.Call(c.handle, uintptr(defID))
			return nil, fmt.Errorf("SimConnect response too short: need %d bytes for %q, got %d total", offset+layouts[i].size, v.Name, len(data))
		}
		if layouts[i].isString {
			// Read null-terminated string from 256-byte block.
			s := data[offset : offset+string256Size]
			end := 0
			for end < len(s) && s[end] != 0 {
				end++
			}
			result[v.Name] = string(s[:end])
		} else {
			bits := binary.LittleEndian.Uint64(data[offset : offset+8])
			result[v.Name] = math.Float64frombits(bits)
		}
		offset += layouts[i].size
	}

	// Clear the definition for reuse.
	procClearDataDefinition.Call(c.handle, uintptr(defID))

	return result, nil
}

func (c *WinClient) SetVariable(name, unit string, value float64) error {
	if err := c.ensureConnected(); err != nil {
		return err
	}

	defID := c.allocDefID()

	namePtr, _ := syscall.BytePtrFromString(name)
	unitPtr, _ := syscall.BytePtrFromString(unit)

	hr, _, _ := procAddToDataDefinition.Call(
		c.handle,
		uintptr(defID),
		uintptr(unsafe.Pointer(namePtr)),
		uintptr(unsafe.Pointer(unitPtr)),
		DataTypeFloat64,
		0,      // epsilon
		Unused, // datum ID: SIMCONNECT_UNUSED
	)
	if hr != HrOK {
		return fmt.Errorf("AddToDataDefinition(%s) failed: 0x%08X", name, hr)
	}

	hr, _, _ = procSetDataOnSimObject.Call(
		c.handle,
		uintptr(defID),
		ObjectIDUser,
		0, // flags
		0, // array count (0 = not an array)
		8, // size of one float64
		uintptr(unsafe.Pointer(&value)),
	)
	if hr != HrOK {
		return fmt.Errorf("SetDataOnSimObject(%s) failed: 0x%08X", name, hr)
	}

	procClearDataDefinition.Call(c.handle, uintptr(defID))
	return nil
}

func (c *WinClient) SendEvent(event string, value int) error {
	if err := c.ensureConnected(); err != nil {
		return err
	}

	c.mu.Lock()
	eventID, ok := c.events[event]
	if !ok {
		eventID = c.nextEvent
		c.nextEvent++

		namePtr, _ := syscall.BytePtrFromString(event)
		hr, _, _ := procMapClientEvent.Call(
			c.handle,
			uintptr(eventID),
			uintptr(unsafe.Pointer(namePtr)),
		)
		if hr != HrOK {
			c.mu.Unlock()
			return fmt.Errorf("MapClientEventToSimEvent(%s) failed: 0x%08X", event, hr)
		}
		c.events[event] = eventID
	}
	c.mu.Unlock()

	hr, _, _ := procTransmitEvent.Call(
		c.handle,
		ObjectIDUser,
		uintptr(eventID),
		uintptr(value),
		GroupPriorityHighest,
		EventFlagGroupIDIsPriority,
	)
	if hr != HrOK {
		return fmt.Errorf("TransmitClientEvent(%s) failed: 0x%08X", event, hr)
	}
	return nil
}

func (c *WinClient) LoadFlightPlan(path string) error {
	if err := c.ensureConnected(); err != nil {
		return err
	}
	pathPtr, _ := syscall.BytePtrFromString(path)
	hr, _, _ := procFlightPlanLoad.Call(c.handle, uintptr(unsafe.Pointer(pathPtr)))
	if hr != HrOK {
		return fmt.Errorf("SimConnect_FlightPlanLoad failed: 0x%08X", hr)
	}
	return nil
}

func (c *WinClient) SaveFlight(path string) error {
	if err := c.ensureConnected(); err != nil {
		return err
	}
	pathPtr, _ := syscall.BytePtrFromString(path)
	hr, _, _ := procFlightSave.Call(c.handle, uintptr(unsafe.Pointer(pathPtr)))
	if hr != HrOK {
		return fmt.Errorf("SimConnect_FlightSave failed: 0x%08X", hr)
	}
	return nil
}

func (c *WinClient) GetPosition() (*Position, error) {
	vars := []VarRequest{
		{Name: "PLANE LATITUDE", Unit: "degrees"},
		{Name: "PLANE LONGITUDE", Unit: "degrees"},
		{Name: "PLANE ALTITUDE", Unit: "feet"},
		{Name: "HEADING INDICATOR", Unit: "degrees"},
		{Name: "AIRSPEED INDICATED", Unit: "knots"},
		{Name: "SIM ON GROUND", Unit: "bool"},
	}
	data, err := c.GetVariables(vars)
	if err != nil {
		return nil, err
	}
	return &Position{
		Latitude:  data["PLANE LATITUDE"].(float64),
		Longitude: data["PLANE LONGITUDE"].(float64),
		Altitude:  data["PLANE ALTITUDE"].(float64),
		Heading:   data["HEADING INDICATOR"].(float64),
		Airspeed:  data["AIRSPEED INDICATED"].(float64),
		OnGround:  data["SIM ON GROUND"].(float64) != 0,
	}, nil
}

func (c *WinClient) SetPosition(pos InitPosition) error {
	if err := c.ensureConnected(); err != nil {
		return err
	}

	defID := c.allocDefID()

	// Define the INITPOSITION struct as the data definition.
	namePtr, _ := syscall.BytePtrFromString("Initial Position")
	unitPtr, _ := syscall.BytePtrFromString("")
	hr, _, _ := procAddToDataDefinition.Call(
		c.handle, uintptr(defID),
		uintptr(unsafe.Pointer(namePtr)),
		uintptr(unsafe.Pointer(unitPtr)),
		DataTypeInitPosition, 0, 0,
	)
	if hr != HrOK {
		return fmt.Errorf("AddToDataDefinition(INITPOSITION) failed: 0x%08X", hr)
	}

	// Marshal the position into the SIMCONNECT_DATA_INITPOSITION struct.
	var buf [InitPositionSize]byte
	binary.LittleEndian.PutUint64(buf[0:8], math.Float64bits(pos.Latitude))
	binary.LittleEndian.PutUint64(buf[8:16], math.Float64bits(pos.Longitude))
	binary.LittleEndian.PutUint64(buf[16:24], math.Float64bits(pos.Altitude))
	binary.LittleEndian.PutUint64(buf[24:32], math.Float64bits(pos.Pitch))
	binary.LittleEndian.PutUint64(buf[32:40], math.Float64bits(pos.Bank))
	binary.LittleEndian.PutUint64(buf[40:48], math.Float64bits(pos.Heading))
	onGround := uint32(0)
	if pos.OnGround {
		onGround = 1
	}
	binary.LittleEndian.PutUint32(buf[48:52], onGround)
	binary.LittleEndian.PutUint64(buf[52:60], math.Float64bits(pos.Airspeed))

	hr, _, _ = procSetDataOnSimObject.Call(
		c.handle,
		uintptr(defID),
		ObjectIDUser,
		0,
		0,
		InitPositionSize,
		uintptr(unsafe.Pointer(&buf[0])),
	)
	if hr != HrOK {
		return fmt.Errorf("SetDataOnSimObject(INITPOSITION) failed: 0x%08X", hr)
	}

	procClearDataDefinition.Call(c.handle, uintptr(defID))
	return nil
}

func (c *WinClient) GetSystemState(stateName string) (string, error) {
	if err := c.ensureConnected(); err != nil {
		return "", err
	}

	reqID := c.allocReqID()
	namePtr, _ := syscall.BytePtrFromString(stateName)

	hr, _, _ := procRequestSystemState.Call(
		c.handle,
		uintptr(reqID),
		uintptr(unsafe.Pointer(namePtr)),
	)
	if hr != HrOK {
		return "", fmt.Errorf("RequestSystemState(%s) failed: 0x%08X", stateName, hr)
	}

	data, err := c.waitForResponse(reqID, 5*time.Second)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (c *WinClient) SetCamera(x, y, z, pitch, bank, heading float64) error {
	if err := c.ensureConnected(); err != nil {
		return err
	}
	hr, _, _ := procCameraSetRelative6DOF.Call(
		c.handle,
		uintptr(math.Float32bits(float32(x))),
		uintptr(math.Float32bits(float32(y))),
		uintptr(math.Float32bits(float32(z))),
		uintptr(math.Float32bits(float32(pitch))),
		uintptr(math.Float32bits(float32(bank))),
		uintptr(math.Float32bits(float32(heading))),
	)
	if hr != HrOK {
		return fmt.Errorf("CameraSetRelative6DOF failed: 0x%08X", hr)
	}
	return nil
}

func (c *WinClient) GetCamera() (*CameraState, error) {
	return nil, fmt.Errorf("GetCamera not yet implemented (requires camera SimVars)")
}

func (c *WinClient) GetAirport(_ string) (*Airport, error) {
	return nil, fmt.Errorf("GetAirport not yet implemented (requires facility data request)")
}

func (c *WinClient) CreateAIAircraft(title, tailNumber string, pos InitPosition) (uint32, error) {
	if err := c.ensureConnected(); err != nil {
		return 0, err
	}

	reqID := c.allocReqID()
	titlePtr, _ := syscall.BytePtrFromString(title)
	tailPtr, _ := syscall.BytePtrFromString(tailNumber)

	// Marshal InitPosition.
	var buf [InitPositionSize]byte
	binary.LittleEndian.PutUint64(buf[0:8], math.Float64bits(pos.Latitude))
	binary.LittleEndian.PutUint64(buf[8:16], math.Float64bits(pos.Longitude))
	binary.LittleEndian.PutUint64(buf[16:24], math.Float64bits(pos.Altitude))
	binary.LittleEndian.PutUint64(buf[24:32], math.Float64bits(pos.Pitch))
	binary.LittleEndian.PutUint64(buf[32:40], math.Float64bits(pos.Bank))
	binary.LittleEndian.PutUint64(buf[40:48], math.Float64bits(pos.Heading))
	onGround := uint32(0)
	if pos.OnGround {
		onGround = 1
	}
	binary.LittleEndian.PutUint32(buf[48:52], onGround)
	binary.LittleEndian.PutUint64(buf[52:60], math.Float64bits(pos.Airspeed))

	hr, _, _ := procAICreateParked.Call(
		c.handle,
		uintptr(unsafe.Pointer(titlePtr)),
		uintptr(unsafe.Pointer(tailPtr)),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(reqID),
	)
	if hr != HrOK {
		return 0, fmt.Errorf("AICreateParkedATCAircraft failed: 0x%08X", hr)
	}

	// The object ID comes back async — for now return the request ID.
	// Full implementation would wait for SIMCONNECT_RECV_ASSIGNED_OBJECT_ID.
	return reqID, nil
}

func (c *WinClient) RemoveAIObject(objectID uint32) error {
	if err := c.ensureConnected(); err != nil {
		return err
	}
	reqID := c.allocReqID()
	hr, _, _ := procAIRemoveObject.Call(
		c.handle,
		uintptr(objectID),
		uintptr(reqID),
	)
	if hr != HrOK {
		return fmt.Errorf("AIRemoveObject failed: 0x%08X", hr)
	}
	return nil
}
