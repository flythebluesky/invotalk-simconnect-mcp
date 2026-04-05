//go:build !windows

package simconnect

import "fmt"

var errNotWindows = fmt.Errorf("SimConnect requires Windows")

// StubClient is a no-op client for non-Windows platforms.
type StubClient struct{}

func NewClient(dllPath string) Client { return &StubClient{} }

func (c *StubClient) Connect() error                          { return errNotWindows }
func (c *StubClient) Close()                                  {}
func (c *StubClient) IsConnected() bool                       { return false }
func (c *StubClient) GetVariables(_ []VarRequest) (map[string]interface{}, error) {
	return nil, errNotWindows
}
func (c *StubClient) SetVariable(_, _ string, _ float64) error { return errNotWindows }
func (c *StubClient) SendEvent(_ string, _ int) error          { return errNotWindows }
func (c *StubClient) LoadFlightPlan(_ string) error            { return errNotWindows }
func (c *StubClient) SaveFlight(_ string) error                { return errNotWindows }
func (c *StubClient) GetPosition() (*Position, error)          { return nil, errNotWindows }
func (c *StubClient) SetPosition(_ InitPosition) error         { return errNotWindows }
func (c *StubClient) GetAirport(_ string) (*Airport, error)    { return nil, errNotWindows }
func (c *StubClient) CreateAIAircraft(_, _ string, _ InitPosition) (uint32, error) {
	return 0, errNotWindows
}
func (c *StubClient) RemoveAIObject(_ uint32) error                        { return errNotWindows }
func (c *StubClient) GetCamera() (*CameraState, error)                     { return nil, errNotWindows }
func (c *StubClient) SetCamera(_, _, _, _, _, _ float64) error             { return errNotWindows }
func (c *StubClient) GetSystemState(_ string) (string, error)              { return "", errNotWindows }
