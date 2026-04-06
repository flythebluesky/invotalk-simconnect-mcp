# Security Policy

## Reporting a Vulnerability

Please do **not** file a public GitHub issue for security vulnerabilities.

Report vulnerabilities privately via [GitHub Security Advisories](https://github.com/flythebluesky/invotalk-simconnect-mcp/security/advisories/new).

We will respond within 7 days and coordinate a fix and disclosure timeline with you.

## Scope

This server runs locally and communicates with MSFS via SimConnect on the same machine.
The primary attack surface is the optional HTTP transport — if you enable it and expose
it on a network, ensure `AUTH_BEARER_TOKENS` is set and TLS is enabled.
