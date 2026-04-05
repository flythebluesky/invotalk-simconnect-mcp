package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
)

func GetAirportTool() mcp.Tool {
	return mcp.NewTool("get_airport",
		mcp.WithDescription("Query airport data by ICAO code. Returns runways, taxiways, parking info."),
		mcp.WithToolAnnotation(mcp.ToolAnnotation{
			Title:        "Get Airport",
			ReadOnlyHint: mcp.ToBoolPtr(true),
		}),
		mcp.WithString("icao", mcp.Required(), mcp.Description("ICAO airport code (e.g. KJFK, EGLL)")),
	)
}

func HandleGetAirport(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	start := time.Now()
	icao, err := request.RequireString("icao")
	if err != nil {
		return mcp.NewToolResultError("icao is required"), nil
	}
	airport, err := SimClient.GetAirport(icao)
	if err != nil {
		slog.ErrorContext(ctx, "tool.get_airport", "status", "error", "icao", icao, "error", err.Error())
		return mcp.NewToolResultError(fmt.Sprintf("SimConnect error: %v", err)), nil
	}
	out, _ := json.Marshal(airport)
	slog.InfoContext(ctx, "tool.get_airport", "status", "ok", "icao", icao, "duration_ms", time.Since(start).Milliseconds())
	return mcp.NewToolResultText(string(out)), nil
}

func GetNavaidsTool() mcp.Tool {
	return mcp.NewTool("get_navaids",
		mcp.WithDescription("Query VORs, NDBs, and waypoints near a position. Not yet implemented."),
		mcp.WithToolAnnotation(mcp.ToolAnnotation{
			Title:        "Get Navaids",
			ReadOnlyHint: mcp.ToBoolPtr(true),
		}),
		mcp.WithNumber("latitude", mcp.Required(), mcp.Description("Center latitude")),
		mcp.WithNumber("longitude", mcp.Required(), mcp.Description("Center longitude")),
		mcp.WithNumber("radius_nm", mcp.Description("Search radius in nautical miles (default 50)")),
	)
}

func HandleGetNavaids(ctx context.Context, _ mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	slog.InfoContext(ctx, "tool.get_navaids", "status", "not_implemented")
	return mcp.NewToolResultError("get_navaids is not yet implemented"), nil
}
