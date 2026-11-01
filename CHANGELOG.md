# Changelog

All notable changes are documented here, in [Keep a Changelog](https://keepachangelog.com/en/1.1.0/) format. This project uses [Semantic Versioning](https://semver.org/).

## [Unreleased]

## [0.1.0] - 2026-11-01
### Added
- `review` command: groups race reports by their two code locations, ranks them by tests hit, and lists lines in more than one race
- `compare` command: races new in one log, gone, or still there, matched by function
- Reads plain, `-v` and `-json` `go test` output, `go run -race` output and GitHub Actions logs
- Hints RS101 to RS108 for maps, mixed atomics, slices, shared standard library values, self-racing lines, closures, test-only races and lost frames
- `--json` and `--fail` options
- Claude Code plugin with the `triaging-go-data-races` skill
