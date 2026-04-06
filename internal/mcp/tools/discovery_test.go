package tools

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
)

func TestHandleListEvents_All(t *testing.T) {
	req := mcp.CallToolRequest{}
	result, err := HandleListEvents(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.IsError {
		t.Fatalf("result.IsError = true")
	}

	var body map[string]interface{}
	if err := json.Unmarshal([]byte(result.Content[0].(mcp.TextContent).Text), &body); err != nil {
		t.Fatalf("decode: %v", err)
	}
	count := int(body["count"].(float64))
	if count == 0 {
		t.Error("expected events, got 0")
	}
}

func TestHandleListEvents_CategoryFilter(t *testing.T) {
	req := mcp.CallToolRequest{}
	req.Params.Arguments = map[string]any{"category": "Gear"}

	result, err := HandleListEvents(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var body map[string]interface{}
	if err := json.Unmarshal([]byte(result.Content[0].(mcp.TextContent).Text), &body); err != nil {
		t.Fatalf("decode: %v", err)
	}
	count := int(body["count"].(float64))
	if count == 0 {
		t.Error("expected Gear events, got 0")
	}
	events := body["events"].([]interface{})
	for _, e := range events {
		ev := e.(map[string]interface{})
		if ev["category"] != "Gear" {
			t.Errorf("event category = %q, want Gear", ev["category"])
		}
	}
}

func TestHandleListEvents_CategoryFilterCaseInsensitive(t *testing.T) {
	req := mcp.CallToolRequest{}
	req.Params.Arguments = map[string]any{"category": "gear"}

	result, _ := HandleListEvents(context.Background(), req)
	var body map[string]interface{}
	json.Unmarshal([]byte(result.Content[0].(mcp.TextContent).Text), &body) //nolint:errcheck
	if int(body["count"].(float64)) == 0 {
		t.Error("category filter should be case-insensitive")
	}
}

func TestHandleListEvents_UnknownCategory(t *testing.T) {
	req := mcp.CallToolRequest{}
	req.Params.Arguments = map[string]any{"category": "NoSuchCategory"}

	result, _ := HandleListEvents(context.Background(), req)
	var body map[string]interface{}
	json.Unmarshal([]byte(result.Content[0].(mcp.TextContent).Text), &body) //nolint:errcheck
	if int(body["count"].(float64)) != 0 {
		t.Error("expected 0 events for unknown category")
	}
}

func TestHandleListVariables_All(t *testing.T) {
	req := mcp.CallToolRequest{}
	result, err := HandleListVariables(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var body map[string]interface{}
	if err := json.Unmarshal([]byte(result.Content[0].(mcp.TextContent).Text), &body); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if int(body["count"].(float64)) == 0 {
		t.Error("expected variables, got 0")
	}
}

func TestHandleListVariables_CategoryFilter(t *testing.T) {
	req := mcp.CallToolRequest{}
	req.Params.Arguments = map[string]any{"category": "Speed"}

	result, _ := HandleListVariables(context.Background(), req)
	var body map[string]interface{}
	json.Unmarshal([]byte(result.Content[0].(mcp.TextContent).Text), &body) //nolint:errcheck

	count := int(body["count"].(float64))
	if count == 0 {
		t.Error("expected Speed variables, got 0")
	}
	vars := body["variables"].([]interface{})
	for _, v := range vars {
		vr := v.(map[string]interface{})
		if vr["category"] != "Speed" {
			t.Errorf("variable category = %q, want Speed", vr["category"])
		}
	}
}
