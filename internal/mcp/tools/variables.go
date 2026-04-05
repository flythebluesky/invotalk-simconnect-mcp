package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"time"

	"github.com/flythebluesky/invotalk-simconnect-mcp/internal/simconnect"
	"github.com/mark3labs/mcp-go/mcp"
)

// SimClient is set by main before the server starts.
var SimClient simconnect.Client

func GetVariablesTool() mcp.Tool {
	return mcp.NewTool("get_variables",
		mcp.WithDescription("Read one or more simulation variables from MSFS. Pass an array of {name, unit} objects. Returns current values."),
		mcp.WithToolAnnotation(mcp.ToolAnnotation{
			Title:        "Get Variables",
			ReadOnlyHint: mcp.ToBoolPtr(true),
		}),
		mcp.WithString("variables", mcp.Required(),
			mcp.Description(`JSON array of variables to read, e.g. [{"name":"PLANE ALTITUDE","unit":"feet"},{"name":"AIRSPEED INDICATED","unit":"knots"}]`)),
	)
}

func HandleGetVariables(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	start := time.Now()

	varsJSON, err := request.RequireString("variables")
	if err != nil {
		return mcp.NewToolResultError("variables parameter is required"), nil
	}

	var vars []simconnect.VarRequest
	if err := json.Unmarshal([]byte(varsJSON), &vars); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("invalid variables JSON: %v", err)), nil
	}

	if len(vars) == 0 {
		return mcp.NewToolResultError("at least one variable is required"), nil
	}

	result, err := SimClient.GetVariables(vars)
	if err != nil {
		slog.ErrorContext(ctx, "tool.get_variables", "status", "error", "error", err.Error())
		return mcp.NewToolResultError(fmt.Sprintf("SimConnect error: %v", err)), nil
	}

	out, _ := json.Marshal(result)
	slog.InfoContext(ctx, "tool.get_variables", "status", "ok", "count", len(vars), "duration_ms", time.Since(start).Milliseconds())
	return mcp.NewToolResultText(string(out)), nil
}

func SetVariableTool() mcp.Tool {
	return mcp.NewTool("set_variable",
		mcp.WithDescription("Write a simulation variable value in MSFS."),
		mcp.WithString("name", mcp.Required(), mcp.Description("SimVar name (e.g. GENERAL ENG THROTTLE LEVER POSITION:1)")),
		mcp.WithString("unit", mcp.Required(), mcp.Description("Unit (e.g. percent, feet, degrees)")),
		mcp.WithNumber("value", mcp.Required(), mcp.Description("Value to set")),
	)
}

func HandleSetVariable(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	start := time.Now()

	name, err := request.RequireString("name")
	if err != nil {
		return mcp.NewToolResultError("name is required"), nil
	}
	unit, err := request.RequireString("unit")
	if err != nil {
		return mcp.NewToolResultError("unit is required"), nil
	}
	args := request.GetArguments()
	value, ok := args["value"].(float64)
	if !ok {
		return mcp.NewToolResultError("value is required and must be a number"), nil
	}

	if err := SimClient.SetVariable(name, unit, value); err != nil {
		slog.ErrorContext(ctx, "tool.set_variable", "status", "error", "error", err.Error())
		return mcp.NewToolResultError(fmt.Sprintf("SimConnect error: %v", err)), nil
	}

	out, _ := json.Marshal(map[string]interface{}{
		"status": "ok",
		"name":   name,
		"unit":   unit,
		"value":  value,
	})
	slog.InfoContext(ctx, "tool.set_variable", "status", "ok", "name", name, "duration_ms", time.Since(start).Milliseconds())
	return mcp.NewToolResultText(string(out)), nil
}
