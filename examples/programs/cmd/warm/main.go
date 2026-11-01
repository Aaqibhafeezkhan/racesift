// Command warm fills the price cache while serving a lookup, like a service
// that starts taking traffic before its cache warm-up finishes.
package main

import (
	"fmt"
	"sync"

	"example.com/shop/inventory"
)

func main() {
	c := inventory.NewPriceCache()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			c.Set(fmt.Sprintf("sku-%d", i), i)
		}
	}()
	p, ok := c.Get("sku-1")
	wg.Wait()
	fmt.Println("sku-1:", p, ok)
}
