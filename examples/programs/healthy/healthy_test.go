package healthy

import (
	"sync"
	"testing"
)

func TestConcurrentUse(t *testing.T) {
	c := NewPriceCache()
	var s Stats
	var a Audit
	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Set("sku-1", 100)
			c.Get("sku-1")
			s.Hit()
			a.Record("sku-1", 1)
		}()
	}
	if _, err := LoadAll(c, []string{"sku-1", "sku-9"}); err == nil {
		t.Fatal("want an error for sku-9")
	}
	wg.Wait()
	if s.Snapshot() != 8 || a.String() == "" {
		t.Fatalf("got %d hits", s.Snapshot())
	}
}
