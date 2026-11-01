# 0001: One standard-library Go file

**Decision:** the tool is a single Go file with no dependencies, run with `go run`, bundled in a skill.

**Why:** anyone with race reports already has Go, so `go run` is the whole install. An agent can already run `go test -race` and a script, so an MCP server would add a binary and a package without giving the agent anything new. One file also runs straight from the plugin folder.

**Alternatives:** an MCP server with the official Go SDK, like `stuckgo`. It would make sense if agents needed to call this without a shell. A Python script would avoid needing Go, but people with Go race reports have Go.
