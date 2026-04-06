package tools

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
)

func TestHandleGetVariables_Success(t *testing.T) {
	mc := &mockClient{
		connected: true,
		getVarsResult: map[string]interface{}{
			"PLANE ALTITUDE":    35000.0,
			"AIRSPEED INDICATED": 450.0,
		},
	}
	SimClient = mc

	varsJSON := `[{"name":"PLANE ALTITUDE","unit":"feet"},{"name":"AIRSPEED INDICATED","unit":"knots"}]`
	req := mcp.CallToolRequest{}
	req.Params.Arguments = map[string]any{"variables": varsJSON}

	result, err := HandleGetVariables(context.Background(), req)
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
	if body["PLANE ALTITUDE"] != 35000.0 {
		t.Errorf("PLANE ALTITUDE = %v, want 35000", body["PLANE ALTITUDE"])
	}
}

func TestHandleGetVariables_MissingParam(t *testing.T) {
	SimClient = &mockClient{}

	req := mcp.CallToolRequest{}

	result, err := HandleGetVariables(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("expected error result for missing variables param")
	}
}

func TestHandleGetVariables_InvalidJSON(t *testing.T) {
	SimClient = &mockClient{}

	req := mcp.CallToolRequest{}
	req.Params.Arguments = map[string]any{"variables": "not-json"}

	result, _ := HandleGetVariables(context.Background(), req)
	if !result.IsError {
		t.Error("expected error result for invalid JSON")
	}
}

func TestHandleGetVariables_EmptyArray(t *testing.T) {
	SimClient = &mockClient{}

	req := mcp.CallToolRequest{}
	req.Params.Arguments = map[string]any{"variables": "[]"}

	result, _ := HandleGetVariables(context.Background(), req)
	if !result.IsError {
		t.Error("expected error result for empty variables array")
	}
}

func TestHandleGetVariables_ClientError(t *testing.T) {
	mc := &mockClient{getVarsErr: errors.New("not connected")}
	SimClient = mc

	req := mcp.CallToolRequest{}
	req.Params.Arguments = map[string]any{"variables": `[{"name":"PLANE ALTITUDE","unit":"feet"}]`}

	result, _ := HandleGetVariables(context.Background(), req)
	if !result.IsError {
		t.Error("expected error result when client returns error")
	}
}

func TestHandleSetVariable_Success(t *testing.T) {
	mc := &mockClient{}
	SimClient = mc

	req := mcp.CallToolRequest{}
	req.Params.Arguments = map[string]any{
		"name":  "GENERAL ENG THROTTLE LEVER POSITION:1",
		"unit":  "percent",
		"value": float64(75),
	}

	result, err := HandleSetVariable(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.IsError {
		t.Fatalf("result.IsError = true")
	}
	if mc.lastSetVarName != "GENERAL ENG THROTTLE LEVER POSITION:1" {
		t.Errorf("lastSetVarName = %q", mc.lastSetVarName)
	}
	if mc.lastSetVarValue != 75 {
		t.Errorf("lastSetVarValue = %v, want 75", mc.lastSetVarValue)
	}
}

func TestHandleSetVariable_MissingValue(t *testing.T) {
	SimClient = &mockClient{}

	req := mcp.CallToolRequest{}
	req.Params.Arguments = map[string]any{"name": "PLANE ALTITUDE", "unit": "feet"}

	result, _ := HandleSetVariable(context.Background(), req)
	if !result.IsError {
		t.Error("expected error result for missing value param")
	}
}

func TestHandleSetVariable_ClientError(t *testing.T) {
	mc := &mockClient{setVarErr: errors.New("read-only variable")}
	SimClient = mc

	req := mcp.CallToolRequest{}
	req.Params.Arguments = map[string]any{"name": "PLANE ALTITUDE", "unit": "feet", "value": float64(0)}

	result, _ := HandleSetVariable(context.Background(), req)
	if !result.IsError {
		t.Error("expected error result when client returns error")
	}
}
