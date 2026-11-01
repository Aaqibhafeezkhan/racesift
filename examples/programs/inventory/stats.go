package inventory

import "sync/atomic"

// Stats counts requests. Hit uses an atomic add, but Snapshot reads the field plainly.
type Stats struct {
	requests int64
}

func (s *Stats) Hit() {
	atomic.AddInt64(&s.requests, 1)
}

func (s *Stats) Snapshot() int64 {
	return s.requests
}
