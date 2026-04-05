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

func CreateAIAircraftTool() mcp.Tool {
	return mcp.NewTool("create_ai_aircraft",
		mcp.WithDescription("Spawn an AI aircraft at a position. Returns an object ID for later removal."),
		mcp.WithString("title", mcp.Required(), mcp.Description("Aircraft title (e.g. Boeing 737-800)")),
		mcp.WithString("tail_number", mcp.Required(), mcp.Description("Tail number (e.g. N12345)")),
		mcp.WithNumber("latitude", mcp.Required(), mcp.Description("Latitude in degrees")),
		mcp.WithNumber("longitude", mcp.Required(), mcp.Description("Longitude in degrees")),
		mcp.WithNumber("altitude", mcp.Required(), mcp.Description("Altitude in feet MSL")),
		mcp.WithNumber("heading", mcp.Required(), mcp.Description("Heading in degrees")),
	)
}

func HandleCreateAIAircraft(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	start := time.Now()
	title, _ := request.RequireString("title")
	tail, _ := request.RequireString("tail_number")
	args := request.GetArguments()
	lat, _ := args["latitude"].(float64)
	lon, _ := args["longitude"].(float64)
	alt, _ := args["altitude"].(float64)
	hdg, _ := args["heading"].(float64)

	pos := simconnect.InitPosition{
		Latitude: lat, Longitude: lon, Altitude: alt, Heading: hdg, OnGround: true,
	}

	objectID, err := SimClient.CreateAIAircraft(title, tail, pos)
	if err != nil {
		slog.ErrorContext(ctx, "tool.create_ai_aircraft", "status", "error", "error", err.Error())
		return mcp.NewToolResultError(fmt.Sprintf("SimConnect error: %v", err)), nil
	}

	out, _ := json.Marshal(map[string]interface{}{"status": "ok", "object_id": objectID})
	slog.InfoContext(ctx, "tool.create_ai_aircraft", "status", "ok", "title", title, "duration_ms", time.Since(start).Milliseconds())
	return mcp.NewToolResultText(string(out)), nil
}

func RemoveAIObjectTool() mcp.Tool {
	return mcp.NewTool("remove_ai_object",
		mcp.WithDescription("Remove a previously created AI object by its ID."),
		mcp.WithNumber("object_id", mcp.Required(), mcp.Description("Object ID returned by create_ai_aircraft")),
	)
}

func HandleRemoveAIObject(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	start := time.Now()
	args := request.GetArguments()
	id, ok := args["object_id"].(float64)
	if !ok {
		return mcp.NewToolResultError("object_id is required"), nil
	}
	if err := SimClient.RemoveAIObject(uint32(id)); err != nil {
		slog.ErrorContext(ctx, "tool.remove_ai_object", "status", "error", "error", err.Error())
		return mcp.NewToolResultError(fmt.Sprintf("SimConnect error: %v", err)), nil
	}
	out, _ := json.Marshal(map[string]string{"status": "ok"})
	slog.InfoContext(ctx, "tool.remove_ai_object", "status", "ok", "object_id", fmt.Sprintf("%d", int(id)), "duration_ms", time.Since(start).Milliseconds())
	return mcp.NewToolResultText(string(out)), nil
}
