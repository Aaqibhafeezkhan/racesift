//go:build !v2

package inventory

import "sync"

// PriceCache caches prices by SKU. Set takes the lock, but Get forgets to.
type PriceCache struct {
	mu     sync.Mutex
	prices map[string]int
}

func NewPriceCache() *PriceCache {
	return &PriceCache{prices: map[string]int{}}
}

func (c *PriceCache) Get(sku string) (int, bool) {
	p, ok := c.prices[sku]
	return p, ok
}

func (c *PriceCache) Set(sku string, price int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.prices[sku] = price
}
