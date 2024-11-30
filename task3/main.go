package main

import (
	"fmt"
	"sync"
)

func GetSquare(out chan int, v int) {
	defer close(out)
	out <- v * v
}

func main() {
	var wg sync.WaitGroup
	numbers := []int{2, 4, 6, 8, 10}

	calculatedSum := 0
	for _, v := range numbers {
		wg.Add(1)

		go func() {
			out := make(chan int)
			go GetSquare(out, v)
			calculatedSum += <-out
			defer wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("Calculations completed. Sum: %v\n", calculatedSum)
}
