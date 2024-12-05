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

	// вместо мьютекса использую инкрементирование переменной через атомик, но это тоже не гарантирует корректной
	// конкурентной инкрементации переменной. чтобы быть уверенным в том, что горутины правильно увеличат переменную к
	// моменту её вывода на экран, использую WaitGroup
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
