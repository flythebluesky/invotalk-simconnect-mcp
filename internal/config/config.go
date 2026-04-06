package config

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Transport     string // "http" or "stdio"
	Port          int
	LogLevel      string
	TLSEnabled    bool
	BearerTokens  []string
	AuthDisabled  bool
	SimConnectDLL string // optional DLL path override
}

func Load() (Config, error) {
	cfg := Config{
		Transport:     envString("TRANSPORT", "stdio"),
		Port:          envInt("PORT", 8443),
		LogLevel:      envString("LOG_LEVEL", "info"),
		TLSEnabled:    !strings.EqualFold(envString("TLS_ENABLED", "true"), "false"),
		SimConnectDLL: envString("SIMCONNECT_DLL_PATH", "SimConnect.dll"),
	}

	tokens := os.Getenv("AUTH_BEARER_TOKENS")
	if strings.EqualFold(tokens, "disabled") {
		cfg.AuthDisabled = true
	} else if tokens != "" {
		for _, t := range strings.Split(tokens, ",") {
			if trimmed := strings.TrimSpace(t); trimmed != "" {
				cfg.BearerTokens = append(cfg.BearerTokens, trimmed)
			}
		}
	} else {
		cfg.AuthDisabled = true
	}

	if cfg.Transport != "http" && cfg.Transport != "stdio" {
		return Config{}, fmt.Errorf("TRANSPORT must be 'http' or 'stdio', got %q", cfg.Transport)
	}

	return cfg, nil
}

func envString(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func envInt(key string, fallback int) int {
	if v := os.Getenv(key); v != "" {
		n, err := strconv.Atoi(v)
		if err != nil {
			slog.Warn("invalid integer env var", "key", key, "value", v, "default", fallback)
			return fallback
		}
		return n
	}
	return fallback
}
