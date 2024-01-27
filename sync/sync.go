package main

import "sync"

/*
When to use locks over channels and goroutines?
- Use channels when passing ownership of data
- Use mutexes for managing state

More here -> https://go.dev/wiki/MutexOrChannel

-----------------------------------------------

NOTE: Remember to use `go vet` in your build scripts as it can alert you to some subtle bugs in your code before they hit your poor users.
*/

type Counter struct {
	mu    sync.Mutex
	value int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
