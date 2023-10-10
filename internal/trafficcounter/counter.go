package trafficcounter

import "sync"

type TrafficCounter struct {
	bytesTransferred int
	mu               sync.Mutex
}

func NewTrafficCounter() *TrafficCounter {
	return &TrafficCounter{}
}

func (tc *TrafficCounter) AddBytes(bytes int) {
	tc.mu.Lock()
	defer tc.mu.Unlock()
	tc.bytesTransferred += bytes
}

func (tc *TrafficCounter) IsExceeded(limit int) bool {
	tc.mu.Lock()
	defer tc.mu.Unlock()
	return tc.bytesTransferred > limit*1024*1024
}
