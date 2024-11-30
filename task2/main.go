package main

import (
	"fmt"
	"sync"
)

func GetSquare(v int) {
	fmt.Printf("Square of %v = %v\n", v, v*v)
}

func main() {
	var wg sync.WaitGroup
	numbers := []int{2, 4, 6, 8, 10}

	for _, v := range numbers {
		wg.Add(1)

		go func() {
			GetSquare(v)
			defer wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Calculations completed")
}
