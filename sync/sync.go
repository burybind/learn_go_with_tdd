package sync

import "sync"

type Counter struct {
	mu    sync.Mutex
	value int
}

// Inc increments the counter
func (c *Counter) Inc() {
	c.mu.Lock() // lock the mutex to be concurrent-safe. Every goroutine that attempts to call this method will have to wait for the mutex to be unlocked
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
