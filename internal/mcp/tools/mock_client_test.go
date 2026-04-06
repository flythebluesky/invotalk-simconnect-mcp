package tools

import (
	"github.com/flythebluesky/invotalk-simconnect-mcp/internal/simconnect"
)

// mockClient is a controllable SimConnect client for testing.
type mockClient struct {
	connected     bool
	getVarsResult map[string]interface{}
	getVarsErr    error
	setVarErr     error
	sendEventErr  error
	getPositionResult *simconnect.Position
	getPositionErr    error
	setPositionErr    error
	getSystemStateResult string
	getSystemStateErr    error
	getCameraResult *simconnect.CameraState
	getCameraErr    error
	setCameraErr    error

	lastSentEvent string
	lastSentValue int
	lastSetVarName  string
	lastSetVarUnit  string
	lastSetVarValue float64
}

func (m *mockClient) Connect() error    { return nil }
func (m *mockClient) Close()            {}
func (m *mockClient) IsConnected() bool { return m.connected }

func (m *mockClient) GetVariables(_ []simconnect.VarRequest) (map[string]interface{}, error) {
	return m.getVarsResult, m.getVarsErr
}

func (m *mockClient) SetVariable(name, unit string, value float64) error {
	m.lastSetVarName = name
	m.lastSetVarUnit = unit
	m.lastSetVarValue = value
	return m.setVarErr
}

func (m *mockClient) SendEvent(event string, value int) error {
	m.lastSentEvent = event
	m.lastSentValue = value
	return m.sendEventErr
}

func (m *mockClient) LoadFlightPlan(_ string) error                             { return nil }
func (m *mockClient) SaveFlight(_ string) error                                 { return nil }
func (m *mockClient) GetPosition() (*simconnect.Position, error)                { return m.getPositionResult, m.getPositionErr }
func (m *mockClient) SetPosition(_ simconnect.InitPosition) error               { return m.setPositionErr }
func (m *mockClient) GetAirport(_ string) (*simconnect.Airport, error)          { return nil, nil }
func (m *mockClient) CreateAIAircraft(_, _ string, _ simconnect.InitPosition) (uint32, error) {
	return 0, nil
}
func (m *mockClient) RemoveAIObject(_ uint32) error { return nil }
func (m *mockClient) GetCamera() (*simconnect.CameraState, error) {
	return m.getCameraResult, m.getCameraErr
}
func (m *mockClient) SetCamera(_, _, _, _, _, _ float64) error { return m.setCameraErr }
func (m *mockClient) GetSystemState(_ string) (string, error) {
	return m.getSystemStateResult, m.getSystemStateErr
}
