Add a hint for this race pattern: $ARGUMENTS

1. Write a small racy program for it in `examples/programs/`, run it with `-race`, and save the output in `examples/logs/`.
2. Add the hint in `hints()` in `plugin/scripts/racesift.go` with the next free `RSnnn` ID. Never reuse an ID.
3. Add a test that it fires on the new log, and that it doesn't fire on `healthy.log` or the other logs.
4. Add it to the README's hint table and to `CHANGELOG.md`, then run `make test`.
