package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
)

func GetCameraTool() mcp.Tool {
	return mcp.NewTool("get_camera",
		mcp.WithDescription("Get current camera position and orientation."),
		mcp.WithToolAnnotation(mcp.ToolAnnotation{
			Title:        "Get Camera",
			ReadOnlyHint: mcp.ToBoolPtr(true),
		}),
	)
}

func HandleGetCamera(ctx context.Context, _ mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	start := time.Now()
	cam, err := SimClient.GetCamera()
	if err != nil {
		slog.ErrorContext(ctx, "tool.get_camera", "status", "error", "error", err.Error())
		return mcp.NewToolResultError(fmt.Sprintf("SimConnect error: %v", err)), nil
	}
	out, _ := json.Marshal(cam)
	slog.InfoContext(ctx, "tool.get_camera", "status", "ok", "duration_ms", time.Since(start).Milliseconds())
	return mcp.NewToolResultText(string(out)), nil
}

func SetCameraTool() mcp.Tool {
	return mcp.NewTool("set_camera",
		mcp.WithDescription("Set camera position using 6 degrees of freedom (relative to aircraft)."),
		mcp.WithNumber("x", mcp.Required(), mcp.Description("X offset (lateral)")),
		mcp.WithNumber("y", mcp.Required(), mcp.Description("Y offset (vertical)")),
		mcp.WithNumber("z", mcp.Required(), mcp.Description("Z offset (forward/back)")),
		mcp.WithNumber("pitch", mcp.Required(), mcp.Description("Pitch angle in degrees")),
		mcp.WithNumber("bank", mcp.Required(), mcp.Description("Bank angle in degrees")),
		mcp.WithNumber("heading", mcp.Required(), mcp.Description("Heading angle in degrees")),
	)
}

func HandleSetCamera(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	start := time.Now()
	args := request.GetArguments()
	x, _ := args["x"].(float64)
	y, _ := args["y"].(float64)
	z, _ := args["z"].(float64)
	pitch, _ := args["pitch"].(float64)
	bank, _ := args["bank"].(float64)
	heading, _ := args["heading"].(float64)

	if err := SimClient.SetCamera(x, y, z, pitch, bank, heading); err != nil {
		slog.ErrorContext(ctx, "tool.set_camera", "status", "error", "error", err.Error())
		return mcp.NewToolResultError(fmt.Sprintf("SimConnect error: %v", err)), nil
	}
	out, _ := json.Marshal(map[string]string{"status": "ok"})
	slog.InfoContext(ctx, "tool.set_camera", "status", "ok", "duration_ms", time.Since(start).Milliseconds())
	return mcp.NewToolResultText(string(out)), nil
}
