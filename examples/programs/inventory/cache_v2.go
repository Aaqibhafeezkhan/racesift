//go:build v2

package inventory

import "sync"

// PriceCache caches prices by SKU. In v2 both methods take the lock.
type PriceCache struct {
	mu     sync.RWMutex
	prices map[string]int
}

func NewPriceCache() *PriceCache {
	return &PriceCache{prices: map[string]int{}}
}

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
