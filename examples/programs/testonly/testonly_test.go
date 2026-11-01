package testonly

import (
	"strings"
	"sync"
	"testing"
)

// The goroutine and the test both write to one strings.Builder.
func TestSummaryLine(t *testing.T) {
	var out strings.Builder
	done := make(chan struct{})
	go func() {
		defer close(done)
		out.WriteString("header\n")
	}()
	out.WriteString("body\n")
	<-done
}

// Every goroutine increments one captured counter.
func TestRetryCount(t *testing.T) {
	attempts := 0
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			attempts++
		}()
	}
	wg.Wait()
	if attempts == 0 {
		t.Fatal("no attempts")
	}
}
