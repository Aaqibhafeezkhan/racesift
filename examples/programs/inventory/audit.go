package inventory

import (
	"bytes"
	"fmt"
)

// Audit collects audit lines in one shared buffer, with no lock.
type Audit struct {
	buf bytes.Buffer
}

func (a *Audit) Record(sku string, qty int) {
	fmt.Fprintf(&a.buf, "reserve %s x%d\n", sku, qty)
}

func (a *Audit) String() string {
	return a.buf.String()
}
