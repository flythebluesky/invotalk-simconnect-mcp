package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
)

func SendEventTool() mcp.Tool {
	return mcp.NewTool("send_event",
		mcp.WithDescription("Send a SimConnect event to MSFS. Any valid SimConnect event name works — use list_events to discover available events."),
		mcp.WithString("event", mcp.Required(), mcp.Description("SimConnect event name (e.g. GEAR_UP, THROTTLE_SET, AUTOPILOT_ON)")),
		mcp.WithNumber("value", mcp.Description("Optional parameter value for events that accept one (e.g. THROTTLE_SET takes 0-16383)")),
	)
}

func HandleSendEvent(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	start := time.Now()

	event, err := request.RequireString("event")
	if err != nil {
		return mcp.NewToolResultError("event is required"), nil
	}

	value := 0
	args := request.GetArguments()
	if v, ok := args["value"].(float64); ok {
		value = int(v)
	}

	if err := SimClient.SendEvent(event, value); err != nil {
		slog.ErrorContext(ctx, "tool.send_event", "status", "error", "event", event, "error", err.Error())
		return mcp.NewToolResultError(fmt.Sprintf("SimConnect error: %v", err)), nil
	}

	out, _ := json.Marshal(map[string]interface{}{
		"status": "ok",
		"event":  event,
		"value":  value,
	})
	slog.InfoContext(ctx, "tool.send_event", "status", "ok", "event", event, "value", value, "duration_ms", time.Since(start).Milliseconds())
	return mcp.NewToolResultText(string(out)), nil
}

func SendEventsTool() mcp.Tool {
	return mcp.NewTool("send_events",
		mcp.WithDescription("Send multiple SimConnect events in sequence with optional delays between them."),
		mcp.WithString("events", mcp.Required(),
			mcp.Description(`JSON array of events, e.g. [{"event":"GEAR_UP"},{"event":"pause","delay_ms":1000},{"event":"FLAPS_UP"}]`)),
	)
}

func HandleSendEvents(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	start := time.Now()

	eventsJSON, err := request.RequireString("events")
	if err != nil {
		return mcp.NewToolResultError("events parameter is required"), nil
	}

	var events []struct {
		Event   string `json:"event"`
		Value   int    `json:"value"`
		DelayMs int    `json:"delay_ms"`
	}
	if err := json.Unmarshal([]byte(eventsJSON), &events); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("invalid events JSON: %v", err)), nil
	}

	for i, e := range events {
		if e.DelayMs > 0 {
			time.Sleep(time.Duration(e.DelayMs) * time.Millisecond)
			continue
		}
		if err := SimClient.SendEvent(e.Event, e.Value); err != nil {
			slog.ErrorContext(ctx, "tool.send_events", "status", "error", "event", e.Event, "index", i, "error", err.Error())
			return mcp.NewToolResultError(fmt.Sprintf("event %d (%s) failed: %v", i, e.Event, err)), nil
		}
	}

	out, _ := json.Marshal(map[string]interface{}{
		"status": "ok",
		"count":  len(events),
	})
	slog.InfoContext(ctx, "tool.send_events", "status", "ok", "count", len(events), "duration_ms", time.Since(start).Milliseconds())
	return mcp.NewToolResultText(string(out)), nil
}
