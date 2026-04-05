package simconnect

// Client provides all SimConnect operations.
type Client interface {
	Connect() error
	Close()
	IsConnected() bool

	// Variables
	GetVariables(vars []VarRequest) (map[string]interface{}, error)
	SetVariable(name, unit string, value float64) error

	// Events
	SendEvent(event string, value int) error

	// Flight plans
	LoadFlightPlan(path string) error
	SaveFlight(path string) error

	// Position
	GetPosition() (*Position, error)
	SetPosition(pos InitPosition) error

	// Facilities
	GetAirport(icao string) (*Airport, error)

	// AI
	CreateAIAircraft(title, tailNumber string, pos InitPosition) (uint32, error)
	RemoveAIObject(objectID uint32) error

	// Camera
	GetCamera() (*CameraState, error)
	SetCamera(x, y, z, pitch, bank, heading float64) error

	// System
	GetSystemState(stateName string) (string, error)
}

// VarRequest specifies a SimVar to read.
type VarRequest struct {
	Name string `json:"name"`
	Unit string `json:"unit"`
}

// Position holds current aircraft position and state.
type Position struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Altitude  float64 `json:"altitude"`
	Heading   float64 `json:"heading"`
	Airspeed  float64 `json:"airspeed"`
	OnGround  bool    `json:"on_ground"`
}

// InitPosition is used to set/teleport aircraft position.
type InitPosition struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Altitude  float64 `json:"altitude"`
	Heading   float64 `json:"heading"`
	Pitch     float64 `json:"pitch"`
	Bank      float64 `json:"bank"`
	Airspeed  float64 `json:"airspeed"`
	OnGround  bool    `json:"on_ground"`
}

// Airport holds facility data for an airport.
type Airport struct {
	ICAO      string   `json:"icao"`
	Name      string   `json:"name"`
	Latitude  float64  `json:"latitude"`
	Longitude float64  `json:"longitude"`
	Altitude  float64  `json:"altitude"`
	Runways   []Runway `json:"runways"`
}

// Runway holds data for a single runway.
type Runway struct {
	ID      string  `json:"id"`
	Heading float64 `json:"heading"`
	Length  float64 `json:"length"`
	Width   float64 `json:"width"`
}

// CameraState holds the current camera position and orientation.
type CameraState struct {
	X       float64 `json:"x"`
	Y       float64 `json:"y"`
	Z       float64 `json:"z"`
	Pitch   float64 `json:"pitch"`
	Bank    float64 `json:"bank"`
	Heading float64 `json:"heading"`
}
