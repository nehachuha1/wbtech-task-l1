package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	counter int64
}

func (c *Counter) IncrementValue() {
	atomic.AddInt64(&c.counter, 1)
}

func main() {
	cntr := &Counter{counter: 0}

	wg := &sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cntr.IncrementValue()
		}()
	}
	wg.Wait()
	fmt.Println(atomic.LoadInt64(&cntr.counter))
}
