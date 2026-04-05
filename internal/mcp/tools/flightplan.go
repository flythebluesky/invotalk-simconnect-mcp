package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
)

func LoadFlightPlanTool() mcp.Tool {
	return mcp.NewTool("load_flight_plan",
		mcp.WithDescription("Load a .pln flight plan file into MSFS. Note: SimConnect provides no confirmation that loading succeeded."),
		mcp.WithString("path", mcp.Required(), mcp.Description("Path to the .pln file")),
	)
}

func HandleLoadFlightPlan(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	start := time.Now()
	path, err := request.RequireString("path")
	if err != nil {
		return mcp.NewToolResultError("path is required"), nil
	}
	if err := SimClient.LoadFlightPlan(path); err != nil {
		slog.ErrorContext(ctx, "tool.load_flight_plan", "status", "error", "error", err.Error())
		return mcp.NewToolResultError(fmt.Sprintf("SimConnect error: %v", err)), nil
	}
	out, _ := json.Marshal(map[string]string{"status": "ok", "path": path})
	slog.InfoContext(ctx, "tool.load_flight_plan", "status", "ok", "path", path, "duration_ms", time.Since(start).Milliseconds())
	return mcp.NewToolResultText(string(out)), nil
}

func SaveFlightTool() mcp.Tool {
	return mcp.NewTool("save_flight",
		mcp.WithDescription("Save the current flight state to a file."),
		mcp.WithString("path", mcp.Required(), mcp.Description("Path to save the flight file")),
	)
}

func HandleSaveFlight(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	start := time.Now()
	path, err := request.RequireString("path")
	if err != nil {
		return mcp.NewToolResultError("path is required"), nil
	}
	if err := SimClient.SaveFlight(path); err != nil {
		slog.ErrorContext(ctx, "tool.save_flight", "status", "error", "error", err.Error())
		return mcp.NewToolResultError(fmt.Sprintf("SimConnect error: %v", err)), nil
	}
	out, _ := json.Marshal(map[string]string{"status": "ok", "path": path})
	slog.InfoContext(ctx, "tool.save_flight", "status", "ok", "path", path, "duration_ms", time.Since(start).Milliseconds())
	return mcp.NewToolResultText(string(out)), nil
}
