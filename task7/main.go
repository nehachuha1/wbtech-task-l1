package main

import (
	"fmt"
	"sync"
)

// конкурентную запись в мапу мы обеспечиваем засчёт единого мьютекса
func main() {
	mu := &sync.RWMutex{}
	counters := map[int]int{}

	for i := 0; i < 10; i++ {
		go func(map[int]int, int) {
			mu.Lock()
			counters[i] = i + 1
			defer mu.Unlock()
		}(counters, i)
	}
	fmt.Printf("%#v", counters)
}
