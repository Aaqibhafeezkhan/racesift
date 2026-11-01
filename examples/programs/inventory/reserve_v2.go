//go:build v2

package inventory

// reserved counts units reserved since start. v2 adds it with no lock.
var reserved int

// Reserve records a reservation and adds it to the running total.
func Reserve(a *Audit, sku string, qty int) int {
	a.Record(sku, qty)
	reserved += qty
	return reserved
}
