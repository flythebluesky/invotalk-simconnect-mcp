package config

import (
	"testing"
)

func TestLoadDefaults(t *testing.T) {
	t.Setenv("TRANSPORT", "stdio")
	t.Setenv("PORT", "")
	t.Setenv("LOG_LEVEL", "")
	t.Setenv("TLS_ENABLED", "")
	t.Setenv("AUTH_BEARER_TOKENS", "disabled")
	t.Setenv("SIMCONNECT_DLL_PATH", "")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Transport != "stdio" {
		t.Errorf("Transport = %q, want stdio", cfg.Transport)
	}
	if cfg.Port != 8443 {
		t.Errorf("Port = %d, want 8443", cfg.Port)
	}
	if cfg.LogLevel != "info" {
		t.Errorf("LogLevel = %q, want info", cfg.LogLevel)
	}
	if !cfg.TLSEnabled {
		t.Error("TLSEnabled = false, want true")
	}
	if cfg.SimConnectDLL != "" {
		t.Errorf("SimConnectDLL = %q, want empty (auto-detect)", cfg.SimConnectDLL)
	}
}

func TestLoadHTTPTransport(t *testing.T) {
	t.Setenv("TRANSPORT", "http")
	t.Setenv("PORT", "9000")
	t.Setenv("TLS_ENABLED", "false")
	t.Setenv("AUTH_BEARER_TOKENS", "disabled")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Transport != "http" {
		t.Errorf("Transport = %q, want http", cfg.Transport)
	}
	if cfg.Port != 9000 {
		t.Errorf("Port = %d, want 9000", cfg.Port)
	}
	if cfg.TLSEnabled {
		t.Error("TLSEnabled = true, want false")
	}
}

func TestLoadInvalidTransport(t *testing.T) {
	t.Setenv("TRANSPORT", "grpc")
	t.Setenv("AUTH_BEARER_TOKENS", "disabled")

	_, err := Load()
	if err == nil {
		t.Fatal("expected error for invalid TRANSPORT, got nil")
	}
}

func TestLoadBearerTokens(t *testing.T) {
	t.Setenv("TRANSPORT", "http")
	t.Setenv("AUTH_BEARER_TOKENS", "token1, token2 , token3")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.AuthDisabled {
		t.Error("AuthDisabled = true, want false")
	}
	if len(cfg.BearerTokens) != 3 {
		t.Errorf("len(BearerTokens) = %d, want 3", len(cfg.BearerTokens))
	}
	if cfg.BearerTokens[0] != "token1" {
		t.Errorf("BearerTokens[0] = %q, want token1", cfg.BearerTokens[0])
	}
}

func TestLoadAuthDisabledWhenEmpty(t *testing.T) {
	t.Setenv("TRANSPORT", "stdio")
	t.Setenv("AUTH_BEARER_TOKENS", "")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !cfg.AuthDisabled {
		t.Error("AuthDisabled = false, want true when AUTH_BEARER_TOKENS is empty")
	}
}

func TestLoadTLSEnabledCaseInsensitive(t *testing.T) {
	t.Setenv("TRANSPORT", "http")
	t.Setenv("TLS_ENABLED", "FALSE")
	t.Setenv("AUTH_BEARER_TOKENS", "disabled")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.TLSEnabled {
		t.Error("TLSEnabled = true, want false for TLS_ENABLED=FALSE")
	}
}

func TestLoadInvalidPort(t *testing.T) {
	t.Setenv("TRANSPORT", "http")
	t.Setenv("PORT", "notanumber")
	t.Setenv("AUTH_BEARER_TOKENS", "disabled")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// Falls back to default on invalid parse
	if cfg.Port != 8443 {
		t.Errorf("Port = %d, want 8443 (default fallback)", cfg.Port)
	}
}
