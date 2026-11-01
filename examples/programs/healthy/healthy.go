// Package healthy does the same work as inventory, with proper locking.
package healthy

import (
	"bytes"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
)

var ErrUnknownSKU = errors.New("unknown sku")

type PriceCache struct {
	mu     sync.RWMutex
	prices map[string]int
}

func NewPriceCache() *PriceCache { return &PriceCache{prices: map[string]int{}} }

func (c *PriceCache) Get(sku string) (int, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	p, ok := c.prices[sku]
	return p, ok
}

func (c *PriceCache) Set(sku string, price int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.prices[sku] = price
}

type Stats struct{ requests atomic.Int64 }

func (s *Stats) Hit()            { s.requests.Add(1) }
func (s *Stats) Snapshot() int64 { return s.requests.Load() }

type Audit struct {
	mu  sync.Mutex
	buf bytes.Buffer
}

func (a *Audit) Record(sku string, qty int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	fmt.Fprintf(&a.buf, "reserve %s x%d\n", sku, qty)
}

func (a *Audit) String() string {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.buf.String()
}

func LoadAll(c *PriceCache, skus []string) ([]int, error) {
	prices := make([]int, len(skus))
	errs := make([]error, len(skus))
	var wg sync.WaitGroup
	for i, sku := range skus {
		wg.Add(1)
		go func() {
			defer wg.Done()
			p, ok := c.Get(sku)
			if !ok {
				errs[i] = ErrUnknownSKU
				return
			}
			prices[i] = p
		}()
	}
	wg.Wait()
	return prices, errors.Join(errs...)
}
