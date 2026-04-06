<p align="center">
  <img src="invotalk_logo.svg" alt="InvoTalk" height="160" />
</p>

# InvoTalk SimConnect MCP Server

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Go](https://img.shields.io/badge/Go-1.24-00ADD8?logo=go&logoColor=white)](https://go.dev/)
[![Go Report Card](https://goreportcard.com/badge/github.com/flythebluesky/invotalk-simconnect-mcp)](https://goreportcard.com/report/github.com/flythebluesky/invotalk-simconnect-mcp)
[![MCP](https://img.shields.io/badge/MCP-compatible-blueviolet)](https://modelcontextprotocol.io/)
[![Platform](https://img.shields.io/badge/platform-Windows-0078D6?logo=windows&logoColor=white)](https://www.microsoft.com/windows)
[![Stars](https://img.shields.io/github/stars/flythebluesky/invotalk-simconnect-mcp?style=social)](https://github.com/flythebluesky/invotalk-simconnect-mcp/stargazers)

MCP server that gives AI agents full control over Microsoft Flight Simulator 2024 via SimConnect.

Read flight data, send cockpit commands, load flight plans, teleport aircraft, spawn AI traffic, and control the camera — all through the [Model Context Protocol](https://modelcontextprotocol.io/).

## Quick Start

### Prerequisites

- Go 1.24+
- SimConnect.dll (from MSFS SDK — place next to the built exe)
- MSFS 2024 running

### Build & Run

```sh
go build -o bin/invotalk-simconnect-mcp ./cmd/server/

# stdio mode (for Claude Code, Cursor, etc.)
TRANSPORT=stdio ./bin/invotalk-simconnect-mcp

# HTTP mode (for remote clients)
TRANSPORT=http PORT=8443 TLS_ENABLED=false ./bin/invotalk-simconnect-mcp
```

### Claude Code Integration

Add to your Claude Code MCP settings:

```json
{
  "mcpServers": {
    "simconnect": {
      "command": "path/to/invotalk-simconnect-mcp",
      "env": {
        "TRANSPORT": "stdio"
      }
    }
  }
}
```

Then ask Claude things like:
- "What's my current altitude and airspeed?"
- "Retract the landing gear"
- "Set heading bug to 270"
- "Teleport me to JFK runway 31L"
- "Run the after takeoff checklist: gear up, flaps up, landing lights off"

## Tools

19 tools across 8 categories:

| Tool | Description |
|------|-------------|
| `list_events` | Discover 328 SimConnect events by category |
| `list_variables` | Discover 242 simulation variables by category |
| `get_variables` | Read SimVars (altitude, speed, fuel, AP state, etc.) |
| `set_variable` | Write a SimVar value |
| `send_event` | Fire a SimConnect event (GEAR_UP, THROTTLE_SET, etc.) |
| `send_events` | Fire multiple events in sequence with delays |
| `get_position` | Aircraft lat/lon/alt/heading/airspeed |
| `set_position` | Teleport aircraft to coordinates |
| `get_autopilot` | Full autopilot state snapshot |
| `load_flight_plan` | Load a .pln file into MSFS |
| `save_flight` | Save current flight to file |
| `get_airport` | Airport data by ICAO code |
| `get_navaids` | Nearby VORs/NDBs/waypoints |
| `create_ai_aircraft` | Spawn AI aircraft |
| `remove_ai_object` | Remove AI object |
| `get_camera` / `set_camera` | Camera position (6DOF) |
| `get_system_state` | Sim running/paused/aircraft loaded |
| `get_connection_status` | SimConnect connection status |

## Event & Variable Coverage

**328 SimConnect events** across 29 categories:

| Category | Events | Highlights |
|----------|--------|------------|
| Autopilot | 49 | Master, altitude, heading, speed, VS, NAV, approach, LNAV/VNAV, FLC, flight director, autothrottle, yaw damper, back course — with ON/OFF/TOGGLE variants |
| Radio/Nav | 24 | COM1/2, NAV1/2, ADF, transponder, altimeter, OBS, DME — set frequencies, swap active/standby |
| Lights | 23 | Landing, strobe, nav, beacon, taxi, cabin, panel, wing, logo, recognition — ON/OFF/TOGGLE for each |
| Engine | 23 | Auto start/shutdown, magnetos, starters, fuel pump, primer, cowl flaps — per-engine for up to 4 engines |
| Views | 19 | Cockpit, external, directional views, zoom, pan |
| Throttle | 16 | Full/cut/set/incr/decr, per-engine, reverse thrust |
| Simulation | 13 | Pause, sim rate, freeze position/altitude/attitude |
| Slew | 12 | Enable/disable, move in all 6 axes, rotate heading |
| Anti-Ice | 12 | Pitot heat, structural, engine, windshield, carb heat — ON/OFF/TOGGLE |
| Mixture | 10 | Rich/lean/set/incr/decr, per-engine, best power |
| Brakes | 10 | Left/right/both, parking brake, autobrakes (disarm through max/RTO) |
| ATC | 10 | Open/close menu, select options 1-8 |
| Trim | 9 | Elevator, rudder, aileron — up/down and SET with value |
| Propeller | 9 | High/low/set/incr/decr, per-engine |
| Fuel | 9 | Selector (all/left/right/off), crossfeed, fuel dump |
| Helicopter | 8 | Rotor brake, governor, clutch, collective |
| Flaps | 8 | Up/down/incr/decr/set, positions 1-3 |
| Failures | 8 | Engine, electrical, pitot, static port, hydraulic, brake, vacuum |
| Electrical | 8 | Battery, alternator, avionics ON/OFF, vacuum pump, alternate static |
| Spoilers | 7 | On/off/toggle/set, arm on/off/toggle |
| Cabin | 7 | Seatbelt sign, no smoking sign, doors, wing fold, tailhook, water rudder |
| Time | 6 | Set zulu hours/minutes/day/year, local clock |
| GPS | 6 | OBS on/off/set, activate waypoint, direct-to, GPS drives NAV1 |
| Controls | 6 | Aileron/elevator/rudder axis set |
| Gear | 5 | Up/down/toggle/set, emergency pump |
| Pressurization | 3 | Cabin altitude set, climb rate set, dump valve |
| Instruments | 3 | Heading gyro set, gyro drift, attitude bars |
| Seaplane | 2 | Water ballast, water rudder |
| Flight Plan | 2 | Activate/deactivate |
| Weight | 1 | Payload station weight set |

**242 simulation variables** across 26 categories:

| Category | Vars | Highlights |
|----------|------|------------|
| Engine | 36 | RPM, N1, N2, throttle position, EGT, ITT, oil pressure/temp, fuel flow, combustion, failure, fire — per-engine for up to 4 engines |
| Radio/Nav | 32 | COM/NAV/ADF frequencies (active + standby), transponder, CDI/GSI needles, glideslope, markers, DME, HSI, TO/FROM |
| Autopilot | 20 | Master, FD, altitude/heading/speed/VS hold + selected values, NAV, approach, FLC, mach, yaw damper, back course, autothrottle |
| Controls | 18 | Flaps position/angle, gear handle + per-wheel extension, spoilers, aileron/elevator/rudder position + trim, deflection angles |
| Aircraft | 14 | Title, ICAO type, model, tail number, airline, flight number, heavy flag, category, V-speeds (Vs0, Vs1, Vc, Vy, Vr, Vto) |
| Weight | 11 | Total/empty/max gross, payload station count + weights (4 stations), CG percent + fwd/aft limits |
| Fuel | 11 | Total quantity/weight, left/right/center tanks, per-engine flow (GPH + PPH), selected quantity, crossfeed, selector |
| GPS | 10 | Ground speed, waypoint distance/ETE/bearing/ID, cross-track error, course to steer, active flight plan, OBS mode/value |
| Environment | 9 | OAT, TAT, ISA temp, wind speed/direction, pressure, visibility, precipitation, sea level pressure |
| Electrical | 9 | Battery, alternator, avionics, pitot heat, total load, voltage, APU switch/generator/RPM/volts |
| Position | 8 | Lat/lon, MSL altitude, AGL, ground elevation, indicated altitude, radio height, pressure altitude |
| Helicopter | 7 | Rotor RPM/percent, brake, collective, governor, clutch |
| Orientation | 6 | True/magnetic heading, heading indicator, pitch, bank, angle of attack |
| Lights | 6 | Landing, strobe, nav, beacon, taxi, cabin |
| Anti-Ice | 6 | Pitot heat, pitot ice %, structural ice %, deice switch, engine anti-ice, windshield deice |
| Speed | 5 | IAS, TAS, ground speed, vertical speed, mach |
| Simulation | 5 | On ground, sim speed, absolute/zulu/local time |
| Pressurization | 4 | Cabin altitude, differential pressure, altitude goal, dump switch |
| Hydraulics | 4 | System 1/2 pressure, pump 1/2 switch |
| Brakes | 4 | Left/right position, parking brake, autobrake setting |
| Warnings | 3 | Stall warning, overspeed warning, AP disconnect |
| Surface | 3 | Surface type, condition, on runway |
| Instruments | 3 | Kohlsman setting, decision height, decision altitude |
| G-Forces | 3 | Current, max, min |
| Doors | 3 | Main door, exit 1, canopy open percent |
| Cabin | 2 | Seatbelt sign, no smoking sign |

## Configuration

Environment variables (see `.env.example`):

| Variable | Default | Description |
|----------|---------|-------------|
| `TRANSPORT` | `stdio` | `stdio` or `http` |
| `PORT` | `8443` | HTTP port |
| `TLS_ENABLED` | `true` | Enable HTTPS |
| `LOG_LEVEL` | `info` | Log verbosity |
| `AUTH_BEARER_TOKENS` | `disabled` | Comma-separated tokens for HTTP auth |
| `SIMCONNECT_DLL_PATH` | `SimConnect.dll` | Custom DLL path |

## Development

Builds on macOS with a stub SimConnect client (returns errors for all operations). The MCP server layer and tool definitions work on any platform.

```sh
make build    # Build binary
make test     # Run tests
make run      # Build and run
```

## Resources

Official MSFS SDK documentation for looking up supported events and variables:

| Resource | Description |
|----------|-------------|
| [SimConnect Event IDs](https://docs.flightsimulator.com/html/Programming_Tools/Event_IDs/Event_IDs.htm) | Full list of all SimConnect events — use these with `send_event` |
| [Simulation Variables](https://docs.flightsimulator.com/html/Programming_Tools/SimVars/Simulation_Variables.htm) | All SimVars — use these with `get_variables` / `set_variable` |
| [Aircraft SimVars](https://docs.flightsimulator.com/html/Programming_Tools/SimVars/Aircraft_SimVars/Aircraft_SimVars.htm) | Aircraft-specific variables (engines, controls, systems) |
| [Environment SimVars](https://docs.flightsimulator.com/html/Programming_Tools/SimVars/Environment_SimVars.htm) | Weather, time, and world variables |
| [Camera SimVars](https://docs.flightsimulator.com/html/Programming_Tools/SimVars/Camera_SimVars.htm) | Camera position and view state |
| [SimConnect SDK Overview](https://docs.flightsimulator.com/html/Programming_Tools/SimConnect/SimConnect_SDK.htm) | Full SimConnect API reference |
| [Model Context Protocol](https://modelcontextprotocol.io/) | MCP specification |

## Related

- [InvoTalk](https://github.com/flythebluesky/invotalk) — Voice command system for MSFS 2024 and iRacing

## License

MIT
