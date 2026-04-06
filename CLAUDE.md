# InvoTalk SimConnect MCP Server

MCP server exposing the Microsoft Flight Simulator 2024 SimConnect API. Lets any MCP client read flight data, send commands, load flight plans, manage AI objects, and control the camera.

## Architecture

```
MCP Client (Claude, etc.) → stdio or HTTP → MCP Server → SimConnect.dll → MSFS 2024
```

- `cmd/server/` — Entry point. Transport selection (stdio/HTTP), config loading, graceful shutdown
- `internal/simconnect/` — SimConnect DLL wrapper via `syscall.NewLazyDLL`. Windows implementation + macOS stub
- `internal/mcp/` — MCP server factory. Registers all tools
- `internal/mcp/tools/` — 19 MCP tool implementations across 8 categories
- `internal/config/` — Environment-based configuration
- `internal/handler/` — HTTP health check endpoint

## System Dependencies

### macOS (development only — stub client)

| Dependency | Install | Purpose |
|------------|---------|---------|
| Go 1.24+ | `brew install go` | Build toolchain |

### Windows (production)

| Dependency | Install | Purpose |
|------------|---------|---------|
| Go 1.24+ | [go.dev/dl](https://go.dev/dl/) | Build toolchain |
| SimConnect.dll | Copy from MSFS SDK or a tool like LorbyAxisAndOhs | MSFS 2024 SimConnect API |

Place `SimConnect.dll` next to the built exe or set `SIMCONNECT_DLL_PATH` env var.

## Build & Run

```sh
# Build
go build -o bin/invotalk-simconnect-mcp ./cmd/server/

# Run in stdio mode (default — for MCP client integration)
TRANSPORT=stdio ./bin/invotalk-simconnect-mcp

# Run in HTTP mode
TRANSPORT=http PORT=8443 TLS_ENABLED=false ./bin/invotalk-simconnect-mcp

# Quick run
make run
```

## Configuration

All via environment variables. See `.env.example`.

| Variable | Default | Description |
|----------|---------|-------------|
| `TRANSPORT` | `stdio` | `stdio` or `http` |
| `PORT` | `8443` | HTTP port (only for http transport) |
| `TLS_ENABLED` | `true` | Enable TLS for HTTP |
| `LOG_LEVEL` | `info` | `debug`, `info`, `warn`, `error` |
| `AUTH_BEARER_TOKENS` | `disabled` | Comma-separated bearer tokens, or `disabled` |
| `SIMCONNECT_DLL_PATH` | *(auto-detect)* | Override DLL path (auto-searches exe dir, MSFS SDK paths) |

## Testing

```sh
go test -race ./...
```

The SimConnect client is a stub on macOS — it returns "SimConnect requires Windows" for all operations. MCP tools and server layer can be tested independently.

## MCP Tools (19 total)

### Discovery
- `list_events` — All known SimConnect events with descriptions (~150 events)
- `list_variables` — All known SimVars with units and descriptions (~100 vars)

### Simulation Variables
- `get_variables` — Read one or more SimVars (altitude, speed, heading, fuel, etc.)
- `set_variable` — Write a SimVar value

### Events
- `send_event` — Fire a SimConnect event (GEAR_UP, THROTTLE_SET, AUTOPILOT_ON, etc.)
- `send_events` — Fire multiple events in sequence with delays

### Position & Autopilot
- `get_position` — Aircraft lat/lon/alt/heading/airspeed/on_ground
- `set_position` — Teleport aircraft
- `get_autopilot` — Full autopilot state snapshot

### Flight Plans
- `load_flight_plan` — Load a .pln file into MSFS
- `save_flight` — Save current flight to file

### Facilities
- `get_airport` — Airport data by ICAO code
- `get_navaids` — VORs/NDBs/waypoints (not yet implemented)

### AI Objects
- `create_ai_aircraft` — Spawn AI aircraft at a position
- `remove_ai_object` — Remove AI object by ID

### Camera & System
- `get_camera` / `set_camera` — Camera position with 6DOF
- `get_system_state` — Sim state (running, paused, aircraft loaded)
- `get_connection_status` — SimConnect connection status

## Key Design Decisions

- **SimConnect via syscall**: Uses `syscall.NewLazyDLL("SimConnect.dll")` — same pattern as the InvoTalk voice command app. No CGo, no third-party Go wrapper.
- **Async dispatch loop**: SimConnect is async. A goroutine polls `GetNextDispatch` and routes responses to waiting callers via channels keyed by request ID.
- **Self-documenting catalogs**: Static catalogs of ~150 events and ~100 SimVars baked into the server. AI clients call `list_events`/`list_variables` to discover what's available before sending commands.
- **macOS stub**: Non-Windows builds compile with a stub client that returns errors. MCP server layer works on any platform.
- **Tool pattern**: Each tool is a definition function + handler function, registered in `internal/mcp/server.go`. Follows the pattern from the `ai-mcp-servers` reference repo.
- **No toggles in sequences**: When using `send_events` for checklists, use SET/ON/OFF events, not TOGGLE — sequences should be idempotent.
