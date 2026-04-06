# Contributing

## Prerequisites

- Go 1.24+
- Windows (required for SimConnect builds; macOS/Linux compile with a stub client)

## Build & Test

```sh
go build ./cmd/server/
go test -race ./...
go vet ./...
```

On non-Windows platforms the SimConnect client is a stub — MCP server logic and tool
definitions can be tested without a running simulator.

## Making Changes

1. Fork the repo and create a branch off `main`
2. Make your changes
3. Run `go vet ./...` and `go test -race ./...`
4. Open a pull request — one approval is required to merge

## Adding Events or SimVars

The event and variable catalogs live in `internal/simconnect/events.go` and
`internal/simconnect/vars.go`. Each entry is a struct with a name, category,
description, and unit. Follow the existing patterns when adding entries.

## Reporting Bugs

Use the [bug report template](.github/ISSUE_TEMPLATE/bug_report.md) and include
the relevant section of `%APPDATA%\Claude\logs\mcp-server-simconnect.log` if
the issue involves the MCP server.

## Questions

Open a [GitHub Discussion](https://github.com/flythebluesky/invotalk-simconnect-mcp/discussions)
or file an issue.
