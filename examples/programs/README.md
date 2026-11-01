# Example programs

A made-up shop service with real data races, used to produce the logs in `../logs/`. It's a separate Go module, so `go test ./...` at the repo root doesn't run these racy tests.

- `inventory/` has the races. Build tag `v2` swaps in a second version, where the price cache race is fixed and a new race is added, for `racesift compare`.
- `healthy/` does the same work with proper locking. Its tests pass with `-race`.
- `cmd/warm/` is a program, so `go run -race` output can be checked too.
- `testonly/` has races only in its tests, to show the hints for test code.

Regenerate the logs from the repo root with `make logs`.
