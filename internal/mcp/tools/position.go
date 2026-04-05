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

func GetPositionTool() mcp.Tool {
	return mcp.NewTool("get_position",
		mcp.WithDescription("Get current aircraft position: latitude, longitude, altitude, heading, airspeed, and whether on ground."),
		mcp.WithToolAnnotation(mcp.ToolAnnotation{
			Title:        "Get Position",
			ReadOnlyHint: mcp.ToBoolPtr(true),
		}),
	)
}

func HandleGetPosition(ctx context.Context, _ mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	start := time.Now()
	pos, err := SimClient.GetPosition()
	if err != nil {
		slog.ErrorContext(ctx, "tool.get_position", "status", "error", "error", err.Error())
		return mcp.NewToolResultError(fmt.Sprintf("SimConnect error: %v", err)), nil
	}
	out, _ := json.Marshal(pos)
	slog.InfoContext(ctx, "tool.get_position", "status", "ok", "duration_ms", time.Since(start).Milliseconds())
	return mcp.NewToolResultText(string(out)), nil
}

func SetPositionTool() mcp.Tool {
	return mcp.NewTool("set_position",
		mcp.WithDescription("Teleport aircraft to a new position."),
		mcp.WithNumber("latitude", mcp.Required(), mcp.Description("Latitude in degrees")),
		mcp.WithNumber("longitude", mcp.Required(), mcp.Description("Longitude in degrees")),
		mcp.WithNumber("altitude", mcp.Required(), mcp.Description("Altitude in feet MSL")),
		mcp.WithNumber("heading", mcp.Required(), mcp.Description("Heading in degrees")),
		mcp.WithNumber("airspeed", mcp.Description("Airspeed in knots (default 0)")),
		mcp.WithBoolean("on_ground", mcp.Description("Whether aircraft is on ground (default false)")),
	)
}

func HandleSetPosition(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	start := time.Now()
	args := request.GetArguments()

	lat, _ := args["latitude"].(float64)
	lon, _ := args["longitude"].(float64)
	alt, _ := args["altitude"].(float64)
	hdg, _ := args["heading"].(float64)
	spd, _ := args["airspeed"].(float64)
	onGround, _ := args["on_ground"].(bool)

	pos := simconnect.InitPosition{
		Latitude:  lat,
		Longitude: lon,
		Altitude:  alt,
		Heading:   hdg,
		Airspeed:  spd,
		OnGround:  onGround,
	}

	if err := SimClient.SetPosition(pos); err != nil {
		slog.ErrorContext(ctx, "tool.set_position", "status", "error", "error", err.Error())
		return mcp.NewToolResultError(fmt.Sprintf("SimConnect error: %v", err)), nil
	}

	out, _ := json.Marshal(map[string]interface{}{"status": "ok", "position": pos})
	slog.InfoContext(ctx, "tool.set_position", "status", "ok", "duration_ms", time.Since(start).Milliseconds())
	return mcp.NewToolResultText(string(out)), nil
}

func GetAutopilotTool() mcp.Tool {
	return mcp.NewTool("get_autopilot",
		mcp.WithDescription("Read all autopilot state: master, flight director, altitude hold, heading hold, nav, approach, VS, speed, LNAV, VNAV, autothrottle."),
		mcp.WithToolAnnotation(mcp.ToolAnnotation{
			Title:        "Get Autopilot State",
			ReadOnlyHint: mcp.ToBoolPtr(true),
		}),
	)
}

func HandleGetAutopilot(ctx context.Context, _ mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	start := time.Now()

	vars := []simconnect.VarRequest{
		{Name: "AUTOPILOT MASTER", Unit: "bool"},
		{Name: "AUTOPILOT FLIGHT DIRECTOR ACTIVE", Unit: "bool"},
		{Name: "AUTOPILOT ALTITUDE LOCK", Unit: "bool"},
		{Name: "AUTOPILOT ALTITUDE LOCK VAR", Unit: "feet"},
		{Name: "AUTOPILOT HEADING LOCK", Unit: "bool"},
		{Name: "AUTOPILOT HEADING LOCK DIR", Unit: "degrees"},
		{Name: "AUTOPILOT AIRSPEED HOLD", Unit: "bool"},
		{Name: "AUTOPILOT AIRSPEED HOLD VAR", Unit: "knots"},
		{Name: "AUTOPILOT VERTICAL HOLD", Unit: "bool"},
		{Name: "AUTOPILOT VERTICAL HOLD VAR", Unit: "feet per minute"},
		{Name: "AUTOPILOT NAV1 LOCK", Unit: "bool"},
		{Name: "AUTOPILOT APPROACH HOLD", Unit: "bool"},
		{Name: "AUTOPILOT THROTTLE ARM", Unit: "bool"},
	}

	result, err := SimClient.GetVariables(vars)
	if err != nil {
		slog.ErrorContext(ctx, "tool.get_autopilot", "status", "error", "error", err.Error())
		return mcp.NewToolResultError(fmt.Sprintf("SimConnect error: %v", err)), nil
	}

	out, _ := json.Marshal(result)
	slog.InfoContext(ctx, "tool.get_autopilot", "status", "ok", "duration_ms", time.Since(start).Milliseconds())
	return mcp.NewToolResultText(string(out)), nil
}
