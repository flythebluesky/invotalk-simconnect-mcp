package tools

import (
	"context"
	"encoding/json"
	"log/slog"
	"strings"
	"time"

	"github.com/flythebluesky/invotalk-simconnect-mcp/internal/simconnect"
	"github.com/mark3labs/mcp-go/mcp"
)

func ListEventsTool() mcp.Tool {
	return mcp.NewTool("list_events",
		mcp.WithDescription("List all known SimConnect events. Returns event names grouped by category with descriptions. Use the event names with send_event."),
		mcp.WithToolAnnotation(mcp.ToolAnnotation{
			Title:        "List Events",
			ReadOnlyHint: mcp.ToBoolPtr(true),
		}),
		mcp.WithString("category", mcp.Description("Filter by category (e.g. Autopilot, Lights, Gear). Omit for all.")),
	)
}

func HandleListEvents(_ context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	start := time.Now()
	args := request.GetArguments()
	category, _ := args["category"].(string)

	var filtered []simconnect.EventInfo
	for _, e := range simconnect.EventCatalog {
		if category == "" || strings.EqualFold(e.Category, category) {
			filtered = append(filtered, e)
		}
	}

	out, _ := json.Marshal(map[string]interface{}{
		"count":  len(filtered),
		"events": filtered,
	})
	slog.Info("tool.list_events", "status", "ok", "count", len(filtered), "duration_ms", time.Since(start).Milliseconds())
	return mcp.NewToolResultText(string(out)), nil
}

func ListVariablesTool() mcp.Tool {
	return mcp.NewTool("list_variables",
		mcp.WithDescription("List all known SimConnect simulation variables. Returns variable names with units and descriptions. Use the names with get_variables."),
		mcp.WithToolAnnotation(mcp.ToolAnnotation{
			Title:        "List Variables",
			ReadOnlyHint: mcp.ToBoolPtr(true),
		}),
		mcp.WithString("category", mcp.Description("Filter by category (e.g. Position, Speed, Autopilot). Omit for all.")),
	)
}

func HandleListVariables(_ context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	start := time.Now()
	args := request.GetArguments()
	category, _ := args["category"].(string)

	var filtered []simconnect.VarInfo
	for _, v := range simconnect.VarCatalog {
		if category == "" || strings.EqualFold(v.Category, category) {
			filtered = append(filtered, v)
		}
	}

	out, _ := json.Marshal(map[string]interface{}{
		"count":     len(filtered),
		"variables": filtered,
	})
	slog.Info("tool.list_variables", "status", "ok", "count", len(filtered), "duration_ms", time.Since(start).Milliseconds())
	return mcp.NewToolResultText(string(out)), nil
}
