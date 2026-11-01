package inventory

import (
	"sync"
	"testing"
)

func fill(c *PriceCache) {
	for i, sku := range []string{"sku-1", "sku-2", "sku-3"} {
		c.Set(sku, 100*(i+1))
	}
}

func TestCacheConcurrentRefresh(t *testing.T) {
	c := NewPriceCache()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() { defer wg.Done(); fill(c) }()
	go func() { defer wg.Done(); c.Get("sku-1") }()
	wg.Wait()
}

func TestLoadAll(t *testing.T) {
	c := NewPriceCache()
	fill(c)
	done := make(chan struct{})
	go func() { defer close(done); fill(c) }()
	prices, err := LoadAll(c, []string{"sku-1", "sku-2", "sku-3", "sku-9"})
	<-done
	if err == nil || len(prices) > 3 {
		t.Fatalf("got %v, %v", prices, err)
	}
}

func TestStats(t *testing.T) {
	var s Stats
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() { defer wg.Done(); s.Hit() }()
	}
	_ = s.Snapshot()
	wg.Wait()
}

func TestReserve(t *testing.T) {
	var a Audit
	var wg sync.WaitGroup
	for _, sku := range []string{"sku-1", "sku-2"} {
		wg.Add(1)
		go func() { defer wg.Done(); Reserve(&a, sku, 1) }()
	}
	wg.Wait()
	if a.String() == "" {
		t.Fatal("no audit lines")
	}
}

func TestCacheWarmAndServe(t *testing.T) {
	c := NewPriceCache()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() { defer wg.Done(); fill(c) }()
	go func() { defer wg.Done(); c.Get("sku-2") }()
	wg.Wait()
}
