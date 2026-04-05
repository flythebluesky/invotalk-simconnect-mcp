package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
)

func GetSystemStateTool() mcp.Tool {
	return mcp.NewTool("get_system_state",
		mcp.WithDescription("Query simulator system state (e.g. sim running, aircraft loaded, flight file)."),
		mcp.WithToolAnnotation(mcp.ToolAnnotation{
			Title:        "Get System State",
			ReadOnlyHint: mcp.ToBoolPtr(true),
		}),
		mcp.WithString("state", mcp.Description("State to query: Sim, AircraftLoaded, FlightLoaded, FlightPlan (default: Sim)")),
	)
}

func HandleGetSystemState(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	start := time.Now()
	args := request.GetArguments()
	state := "Sim"
	if s, ok := args["state"].(string); ok && s != "" {
		state = s
	}

	result, err := SimClient.GetSystemState(state)
	if err != nil {
		slog.ErrorContext(ctx, "tool.get_system_state", "status", "error", "state", state, "error", err.Error())
		return mcp.NewToolResultError(fmt.Sprintf("SimConnect error: %v", err)), nil
	}

	out, _ := json.Marshal(map[string]string{"state": state, "value": result})
	slog.InfoContext(ctx, "tool.get_system_state", "status", "ok", "state", state, "duration_ms", time.Since(start).Milliseconds())
	return mcp.NewToolResultText(string(out)), nil
}

func GetConnectionStatusTool() mcp.Tool {
	return mcp.NewTool("get_connection_status",
		mcp.WithDescription("Check if the MCP server is connected to MSFS via SimConnect."),
		mcp.WithToolAnnotation(mcp.ToolAnnotation{
			Title:        "Connection Status",
			ReadOnlyHint: mcp.ToBoolPtr(true),
		}),
	)
}

func HandleGetConnectionStatus(ctx context.Context, _ mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	connected := SimClient.IsConnected()
	out, _ := json.Marshal(map[string]bool{"connected": connected})
	slog.InfoContext(ctx, "tool.get_connection_status", "connected", connected)
	return mcp.NewToolResultText(string(out)), nil
}
