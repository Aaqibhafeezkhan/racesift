# racesift plugin

This plugin adds the `triaging-go-data-races` skill to Claude Code. When you share `go test -race` output or a CI log with `WARNING: DATA RACE` reports, Claude runs the bundled `scripts/racesift.go` with `go run`. It groups repeated reports into distinct races, ranks them by how many tests hit them, and shows both code locations. Claude then reads the code for each race and proposes a fix. It needs Go 1.22 or newer. See the repository README for examples and limits.
