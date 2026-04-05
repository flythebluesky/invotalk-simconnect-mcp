package main

import (
	"context"
	"crypto/subtle"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/flythebluesky/invotalk-simconnect-mcp/internal/config"
	"github.com/flythebluesky/invotalk-simconnect-mcp/internal/handler"
	mcpserver "github.com/flythebluesky/invotalk-simconnect-mcp/internal/mcp"
	"github.com/flythebluesky/invotalk-simconnect-mcp/internal/mcp/tools"
	"github.com/flythebluesky/invotalk-simconnect-mcp/internal/simconnect"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// Load config.
	cfg, err := config.Load()
	if err != nil {
		slog.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	// Setup structured logging.
	var level slog.Level
	switch strings.ToLower(cfg.LogLevel) {
	case "debug":
		level = slog.LevelDebug
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}
	// In stdio mode stdout is the MCP channel — logs must not corrupt it.
	logWriter := os.Stderr
	if cfg.Transport == "stdio" {
		logPath := os.Getenv("APPDATA") + `\Claude\logs\mcp-server-simconnect.log`
		if f, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644); err != nil {
			fmt.Fprintf(os.Stderr, "warn: could not open log file %s: %v\n", logPath, err)
		} else {
			logWriter = f
		}
	}
	logger := slog.New(slog.NewJSONHandler(logWriter, &slog.HandlerOptions{Level: level}))
	slog.SetDefault(logger)

	// Initialize SimConnect client.
	client := simconnect.NewClient(cfg.SimConnectDLL)
	if err := client.Connect(); err != nil {
		slog.Warn("SimConnect not available (MSFS not running?)", "error", err.Error())
	} else {
		slog.Info("SimConnect connected to MSFS")
	}
	defer client.Close()

	// Wire client into MCP tools.
	tools.SimClient = client

	// Create MCP server.
	mcpSrv := mcpserver.NewMCPServer()

	// Context for graceful shutdown.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if cfg.Transport == "stdio" {
		slog.Info("starting MCP server", "transport", "stdio")
		stdio := server.NewStdioServer(mcpSrv)
		if err := stdio.Listen(ctx, os.Stdin, os.Stdout); err != nil {
			slog.Error("stdio server error", "error", err)
			os.Exit(1)
		}
		return
	}

	// HTTP transport.
	httpMCP := server.NewStreamableHTTPServer(mcpSrv)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(requestLogger)

	r.Get("/health", handler.Health(client.IsConnected))

	r.Group(func(r chi.Router) {
		if !cfg.AuthDisabled {
			r.Use(bearerAuth(cfg.BearerTokens))
		}
		r.Handle("/mcp", httpMCP)
	})

	addr := fmt.Sprintf(":%d", cfg.Port)
	srv := &http.Server{
		Addr:         addr,
		Handler:      r,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 120 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		slog.Info("starting MCP server", "transport", "http", "addr", addr, "tls", cfg.TLSEnabled)
		var listenErr error
		if cfg.TLSEnabled {
			listenErr = srv.ListenAndServeTLS("", "")
		} else {
			listenErr = srv.ListenAndServe()
		}
		if listenErr != nil && listenErr != http.ErrServerClosed {
			slog.Error("server error", "error", listenErr)
			os.Exit(1)
		}
	}()

	<-ctx.Done()
	slog.Info("shutting down...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		slog.Error("shutdown error", "error", err)
	}
}

func requestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		next.ServeHTTP(ww, r)
		if r.URL.Path == "/health" {
			return
		}
		slog.Info("request",
			"method", r.Method,
			"path", r.URL.Path,
			"status", ww.Status(),
			"duration_ms", time.Since(start).Milliseconds(),
			"bytes", ww.BytesWritten(),
		)
	})
}

func bearerAuth(tokens []string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("Authorization")
			if !strings.HasPrefix(auth, "Bearer ") {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}
			token := strings.TrimPrefix(auth, "Bearer ")
			for _, valid := range tokens {
				if subtle.ConstantTimeCompare([]byte(token), []byte(valid)) == 1 {
					next.ServeHTTP(w, r)
					return
				}
			}
			http.Error(w, "unauthorized", http.StatusUnauthorized)
		})
	}
}
