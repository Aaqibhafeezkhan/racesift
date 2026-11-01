package inventory

import (
	"errors"
	"sync"
)

var ErrUnknownSKU = errors.New("unknown sku")

// LoadAll looks up every SKU in parallel. The goroutines append to one slice
// and set one error variable, with no lock.
func LoadAll(c *PriceCache, skus []string) ([]int, error) {
	var prices []int
	var firstErr error
	var wg sync.WaitGroup
	for _, sku := range skus {
		wg.Add(1)
		go func() {
			defer wg.Done()
			p, ok := c.Get(sku)
			if !ok {
				firstErr = ErrUnknownSKU
				return
			}
			prices = append(prices, p)
		}()
	}
	wg.Wait()
	return prices, firstErr
}
