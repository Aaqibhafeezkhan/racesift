# racesift

Claude Code skill and Go CLI that groups and ranks Go race detector reports from `go test -race` output.

## Architecture
`plugin/scripts/racesift.go` holds all the logic in one standard-library file, so `go run` needs no setup. `readLines` unpacks `-json` events and strips CI timestamps, `parseStream` reads each `WARNING: DATA RACE` block and the test it belongs to, `Review` groups reports by their two code locations and adds hints, and `Compare` matches races between two logs. `plugin/skills/triaging-go-data-races/SKILL.md` tells Claude how to use it. See `docs/architecture.md`.

## Commands
- `make test`: gofmt, vet, unit tests and integration tests in `tests/`
- `make lint`: plugin and marketplace validation
- `make demo`: review and compare the example logs
- `make logs`: rebuild the shop logs from `examples/programs`

## Conventions
- Standard library only. Never use the network, and never run the user's code from the tool. The skill may run `go test -race` when the user asks.
- Each hint has a stable `RSnnn` ID. Don't reuse IDs.
- Every hint needs real race output in `examples/logs/` that it fires on, and must stay quiet on the healthy logs.
- Logs in `examples/logs/` are real output. Say so in the file's place in `examples/logs/README.md` if one isn't.

## Do and don't
- Do call hints leads. They come from stack shape, not from the code.
- Don't say code is race free because a log has no reports.

## Definition of done
`make test` and `make lint` pass, and `ROADMAP.md` and docs are updated when relevant.
