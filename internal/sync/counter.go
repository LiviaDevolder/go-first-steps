package sync

import "sync"

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increase() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}

func NewCounter() *Counter {
	return &Counter{}
}
