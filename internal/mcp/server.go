package mcp

import (
	"github.com/flythebluesky/invotalk-simconnect-mcp/internal/mcp/tools"
	"github.com/mark3labs/mcp-go/server"
)

func NewMCPServer() *server.MCPServer {
	s := server.NewMCPServer(
		"invotalk-simconnect-mcp",
		"1.0.0",
		server.WithToolCapabilities(true),
		server.WithInstructions(`SimConnect MCP server for Microsoft Flight Simulator 2024.

Use list_events and list_variables to discover available SimConnect events and simulation variables.
Use send_event to fire events (e.g. GEAR_UP, AUTOPILOT_ON, THROTTLE_SET).
Use get_variables to read simulation state (altitude, speed, heading, fuel, etc.).
Use get_position and get_autopilot for quick aircraft state snapshots.
Use set_position to teleport the aircraft.
Use load_flight_plan to load .pln files into the sim.`),
	)

	// Discovery
	s.AddTool(tools.ListEventsTool(), tools.HandleListEvents)
	s.AddTool(tools.ListVariablesTool(), tools.HandleListVariables)

	// Variables
	s.AddTool(tools.GetVariablesTool(), tools.HandleGetVariables)
	s.AddTool(tools.SetVariableTool(), tools.HandleSetVariable)

	// Events
	s.AddTool(tools.SendEventTool(), tools.HandleSendEvent)
	s.AddTool(tools.SendEventsTool(), tools.HandleSendEvents)

	// Position
	s.AddTool(tools.GetPositionTool(), tools.HandleGetPosition)
	s.AddTool(tools.SetPositionTool(), tools.HandleSetPosition)
	s.AddTool(tools.GetAutopilotTool(), tools.HandleGetAutopilot)

	// Flight Plans
	s.AddTool(tools.LoadFlightPlanTool(), tools.HandleLoadFlightPlan)
	s.AddTool(tools.SaveFlightTool(), tools.HandleSaveFlight)

	// Facilities
	s.AddTool(tools.GetAirportTool(), tools.HandleGetAirport)
	s.AddTool(tools.GetNavaidsTool(), tools.HandleGetNavaids)

	// AI
	s.AddTool(tools.CreateAIAircraftTool(), tools.HandleCreateAIAircraft)
	s.AddTool(tools.RemoveAIObjectTool(), tools.HandleRemoveAIObject)

	// Camera
	s.AddTool(tools.GetCameraTool(), tools.HandleGetCamera)
	s.AddTool(tools.SetCameraTool(), tools.HandleSetCamera)

	// System
	s.AddTool(tools.GetSystemStateTool(), tools.HandleGetSystemState)
	s.AddTool(tools.GetConnectionStatusTool(), tools.HandleGetConnectionStatus)

	return s
}
