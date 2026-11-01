//go:build !v2

package inventory

// Reserve records a reservation. In v1 it only writes the audit line.
func Reserve(a *Audit, sku string, qty int) int {
	a.Record(sku, qty)
	return qty
}
