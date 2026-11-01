---
name: triaging-go-data-races
description: Groups and ranks Go data race reports from `go test -race` or `go run -race` output, then finds the shared variable and missing synchronization for each one. Use when the user has "WARNING: DATA RACE" output, a CI log where tests failed with "race detected during execution of test", asks why the race detector fails, or wants to know if a branch adds new data races.
---

# Triaging Go data races

The race detector prints one long report per race it sees, and the same race often repeats across tests. Group them first with the bundled tool. It only reads text and needs Go installed.

## Get the output

If the user has a log or CI output, save it to a file and use that. Otherwise run the tests, narrowed to the package in question when you know it:

    go test -race ./... 2>&1 | tee /tmp/race.log

`go test -json` output and GitHub Actions logs, with their timestamps, also work. `-race` needs cgo on some platforms. If it fails to build, say so rather than guessing.

## Run the tool

    go run "${CLAUDE_PLUGIN_ROOT}/scripts/racesift.go" review /tmp/race.log

It prints each distinct race once, ranked by how many tests hit it, with both code locations, the tests, where each goroutine started, and hints (`RS101` to `RS108`). "Lines in more than one race" lists code that shows up in several races. Start there, because one missing lock often causes several races. Add `--json` for structured output.

To check a branch, run the tests on both and compare. Races are matched by the functions involved, so line changes elsewhere don't make an old race look new:

    go run "${CLAUDE_PLUGIN_ROOT}/scripts/racesift.go" compare base.log head.log

## Check each race in the code

Hints are leads from the stack shape, not proof. For each race:

1. Read both locations. Name the variable or field they share, and say which goroutines touch it.
2. Find the missing happens-before link: a lock one side forgets, a field read before `wg.Wait`, a value captured by a closure, or a struct replaced while in use.
3. When a side has no frame from the user's code (`RS108`), use where its goroutine started and the other side to find it. Say that the location is inferred.
4. When every frame is in `_test.go` files (`RS107`), check whether the test itself is wrong before changing the code under test.

## Report

1. List races worst first: those hit by the most tests, or on hot lines. Give `file:line` for both sides and the shared variable.
2. Propose the smallest fix in the project's style: a mutex the type already has, an `atomic` type, a channel, or a copy before starting the goroutine.
3. After a fix, suggest rerunning with `go test -race -count=5` on that package. One clean run doesn't prove a race is gone.

## Don't

- Don't edit code unless the user asks.
- Don't suggest `time.Sleep`, retries, or skipping `-race` to make a race go away.
- Don't say code is race free because nothing was reported. The race detector only sees races that happen during that run.
