package tools

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
)

func TestHandleSendEvent_Success(t *testing.T) {
	mc := &mockClient{}
	SimClient = mc

	req := mcp.CallToolRequest{}
	req.Params.Arguments = map[string]any{"event": "GEAR_UP"}

	result, err := HandleSendEvent(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.IsError {
		t.Fatalf("result.IsError = true: %v", result.Content)
	}
	if mc.lastSentEvent != "GEAR_UP" {
		t.Errorf("lastSentEvent = %q, want GEAR_UP", mc.lastSentEvent)
	}
}

func TestHandleSendEvent_WithValue(t *testing.T) {
	mc := &mockClient{}
	SimClient = mc

	req := mcp.CallToolRequest{}
	req.Params.Arguments = map[string]any{"event": "THROTTLE_SET", "value": float64(8192)}

	HandleSendEvent(context.Background(), req) //nolint:errcheck
	if mc.lastSentValue != 8192 {
		t.Errorf("lastSentValue = %d, want 8192", mc.lastSentValue)
	}
}

func TestHandleSendEvent_MissingEvent(t *testing.T) {
	mc := &mockClient{}
	SimClient = mc

	req := mcp.CallToolRequest{}

	result, err := HandleSendEvent(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("expected error result for missing event param")
	}
}

func TestHandleSendEvent_ClientError(t *testing.T) {
	mc := &mockClient{sendEventErr: errors.New("sim not running")}
	SimClient = mc

	req := mcp.CallToolRequest{}
	req.Params.Arguments = map[string]any{"event": "GEAR_UP"}

	result, _ := HandleSendEvent(context.Background(), req)
	if !result.IsError {
		t.Error("expected error result when client returns error")
	}
}

func TestHandleSendEvents_Success(t *testing.T) {
	mc := &mockClient{}
	SimClient = mc

	eventsJSON := `[{"event":"GEAR_UP"},{"event":"FLAPS_UP"}]`
	req := mcp.CallToolRequest{}
	req.Params.Arguments = map[string]any{"events": eventsJSON}

	result, err := HandleSendEvents(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.IsError {
		t.Fatalf("result.IsError = true")
	}

	var body map[string]interface{}
	json.Unmarshal([]byte(result.Content[0].(mcp.TextContent).Text), &body) //nolint:errcheck
	if int(body["count"].(float64)) != 2 {
		t.Errorf("count = %v, want 2", body["count"])
	}
}

func TestHandleSendEvents_InvalidJSON(t *testing.T) {
	mc := &mockClient{}
	SimClient = mc

	req := mcp.CallToolRequest{}
	req.Params.Arguments = map[string]any{"events": "not-json"}

	result, _ := HandleSendEvents(context.Background(), req)
	if !result.IsError {
		t.Error("expected error result for invalid JSON")
	}
}

func TestHandleSendEvents_PartialFailure(t *testing.T) {
	mc := &mockClient{sendEventErr: errors.New("event failed")}
	SimClient = mc

	eventsJSON := `[{"event":"GEAR_UP"},{"event":"FLAPS_UP"}]`
	req := mcp.CallToolRequest{}
	req.Params.Arguments = map[string]any{"events": eventsJSON}

	result, _ := HandleSendEvents(context.Background(), req)
	if !result.IsError {
		t.Error("expected error result when a sequence event fails")
	}
}
