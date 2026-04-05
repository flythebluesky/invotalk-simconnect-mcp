.PHONY: build test lint run dev clean restart-claude

APP_NAME := invotalk-simconnect-mcp.exe
BUILD_DIR := ./

CLAUDE_EXE ?= claude

build:
	go build -o $(BUILD_DIR)/$(APP_NAME) ./cmd/server

test:
	go test -v -race ./...

lint:
	golangci-lint run ./...

# Build and restart Claude Desktop (Claude launches the exe via stdio automatically)
run: build restart-claude

# Run HTTP server in foreground for debugging
dev:
	TRANSPORT=http PORT=9475 TLS_ENABLED=false go run ./cmd/server

restart-claude:
	powershell.exe -Command "Stop-Process -Name invotalk-simconnect-mcp -Force -ErrorAction SilentlyContinue; Stop-Process -Name claude -Force -ErrorAction SilentlyContinue; Start-Sleep -Milliseconds 500; Start-Process '$(CLAUDE_EXE)'"

clean:
	rm -f $(BUILD_DIR)/$(APP_NAME)
	go clean -cache
