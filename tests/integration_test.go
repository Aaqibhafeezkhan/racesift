// Package tests runs racesift on fresh race detector output from the
// programs in examples/programs, the way the skill uses it.
package tests

import (
	"bytes"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

type summary struct {
	Reports int
	Races   []struct {
		Locations []struct {
			File string
			Line int
		}
		Tests []string
		Hints []struct{ ID string }
	}
}

var bin string

func TestMain(m *testing.M) {
	dir, err := os.MkdirTemp("", "racesift")
	if err != nil {
		panic(err)
	}
	bin = filepath.Join(dir, "racesift")
	// Build the single file, as the skill runs it.
	out, err := exec.Command("go", "build", "-o", bin, "../plugin/scripts/racesift.go").CombinedOutput()
	if err != nil {
		panic(string(out))
	}
	code := m.Run()
	os.RemoveAll(dir)
	os.Exit(code)
}

// raceRun runs a go command in examples/programs and returns its output.
// The command fails when races are found, so only the output matters.
func raceRun(t *testing.T, args ...string) []byte {
	t.Helper()
	cmd := exec.Command("go", args...)
	cmd.Dir = "../examples/programs"
	out, _ := cmd.CombinedOutput()
	return out
}

func sift(t *testing.T, log []byte, args ...string) (summary, int) {
	t.Helper()
	cmd := exec.Command(bin, append([]string{"review", "--json", "--fail"}, args...)...)
	cmd.Stdin = bytes.NewReader(log)
	out, err := cmd.Output()
	code := 0
	if e, ok := err.(*exec.ExitError); ok {
		code = e.ExitCode()
	} else if err != nil {
		t.Fatal(err)
	}
	var s summary
	if err := json.Unmarshal(out, &s); err != nil {
		t.Fatalf("bad JSON: %v\n%s", err, out)
	}
	return s, code
}

func hasCacheRace(s summary) bool {
	for _, r := range s.Races {
		lines := map[string]bool{}
		for _, l := range r.Locations {
			lines[filepath.Base(l.File)+":"+strconv.Itoa(l.Line)] = true
		}
		if lines["cache.go:18"] && lines["cache.go:25"] && len(r.Hints) > 0 && r.Hints[0].ID == "RS101" {
			return true
		}
	}
	return false
}

func TestFreshRaceRuns(t *testing.T) {
	if testing.Short() {
		t.Skip("runs go test -race on the examples")
	}
	for _, form := range [][]string{
		{"test", "-race", "-count=1", "./inventory/"},
		{"test", "-race", "-count=1", "-v", "./inventory/"},
		{"test", "-race", "-count=1", "-json", "./inventory/"},
		{"run", "-race", "./cmd/warm"},
	} {
		t.Run(strings.Join(form, " "), func(t *testing.T) {
			log := raceRun(t, form...)
			if !bytes.Contains(log, []byte("WARNING: DATA RACE")) {
				t.Skipf("the race detector didn't report anything here:\n%s", log)
			}
			s, code := sift(t, log)
			if code != 1 {
				t.Errorf("exit %d with races and --fail", code)
			}
			if !hasCacheRace(s) {
				t.Errorf("cache race not found in %d races", len(s.Races))
			}
			if len(s.Races) >= s.Reports && s.Reports > 1 {
				t.Errorf("no grouping: %d races from %d reports", len(s.Races), s.Reports)
			}
		})
	}
}

func TestFreshHealthyRun(t *testing.T) {
	if testing.Short() {
		t.Skip("runs go test -race on the examples")
	}
	log := raceRun(t, "test", "-race", "-count=1", "./healthy/")
	if !bytes.HasPrefix(log, []byte("ok")) {
		t.Fatalf("healthy package failed:\n%s", log)
	}
	s, code := sift(t, log)
	if code != 0 || len(s.Races) != 0 {
		t.Errorf("healthy run: exit %d, %d races", code, len(s.Races))
	}
}
