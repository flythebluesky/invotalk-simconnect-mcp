package simconnect

// HRESULT codes
const (
	HrOK    = 0 // S_OK
	HrError = 1 // E_FAIL
)

// Object IDs
const (
	ObjectIDUser = 0          // SIMCONNECT_OBJECT_ID_USER
	Unused       = 0xFFFFFFFF // SIMCONNECT_UNUSED
)

// Group/Priority
const (
	GroupPriorityHighest       = 1    // SIMCONNECT_GROUP_PRIORITY_HIGHEST
	EventFlagGroupIDIsPriority = 0x10 // SIMCONNECT_EVENT_FLAG_GROUPID_IS_PRIORITY
)

// Data types (SIMCONNECT_DATATYPE)
const (
	DataTypeInvalid  = 0
	DataTypeInt32    = 1
	DataTypeInt64    = 2
	DataTypeFloat32  = 3
	DataTypeFloat64  = 4
	DataTypeString8  = 5
	DataTypeString32 = 6
	DataTypeString64 = 7
	DataTypeString128 = 8
	DataTypeString256 = 9
	DataTypeString260 = 10
	DataTypeStringV  = 11
	DataTypeInitPosition = 12
	DataTypeMarkerState  = 13
	DataTypeWaypoint     = 14
	DataTypeLatLonAlt    = 15
	DataTypeXYZ          = 16
)

// Object types (SIMCONNECT_SIMOBJECT_TYPE)
const (
	SimObjectTypeUser       = 0
	SimObjectTypeAll        = 1
	SimObjectTypeAircraft   = 2
	SimObjectTypeHelicopter = 3
	SimObjectTypeBoat       = 4
	SimObjectTypeGround     = 5
)

// Request period (SIMCONNECT_PERIOD)
const (
	PeriodNever       = 0
	PeriodOnce        = 1
	PeriodVisualFrame = 2
	PeriodSimFrame    = 3
	PeriodSecond      = 4
)

// Recv message types (SIMCONNECT_RECV_ID)
const (
	RecvIDNull                  = 0
	RecvIDException             = 1
	RecvIDOpen                  = 2
	RecvIDQuit                  = 3
	RecvIDEvent                 = 4
	RecvIDSimObjectData         = 8  // response to RequestDataOnSimObject
	RecvIDSimObjectDataByType   = 9  // response to RequestDataOnSimObjectType
	RecvIDSystemState           = 15
	RecvIDEventFilename         = 19
)

// SimConnect INITPOSITION struct layout for SetDataOnSimObject.
// 64 bytes: 7 float64 fields (56 bytes) + 1 uint32 on_ground (4 bytes) + 4 pad.
const InitPositionSize = 64

// VarInfo describes a simulation variable in the catalog.
type VarInfo struct {
	Name        string `json:"name"`
	Unit        string `json:"unit"`
	Description string `json:"description"`
	Category    string `json:"category"`
	ReadOnly    bool   `json:"read_only"`
}

// EventInfo describes a SimConnect event in the catalog.
type EventInfo struct {
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description"`
	HasParam    bool   `json:"has_param"`
}
