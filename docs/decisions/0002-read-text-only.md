# 0002: Read race output as text

**Decision:** read the output of `go test -race` or `go run -race` and never run tests from the tool.

**Why:** the output often comes from CI, where the race happened, and rerunning may not reproduce it, because the race detector only sees races that happen in a run. Reading text also means the tool can't change anything. The skill runs tests only when the user asks.

**Alternatives:** wrapping `go test -race` would save a step, but would hide the real command, its flags and its exit code from the user.
