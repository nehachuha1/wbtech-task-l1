package main

import (
	"fmt"
	"sync"
)

// две горутины, конкурентно отлавливают значения с каналов, третья горутина закидывает числа в один из каналов
// в конце закрываем оба канала
func main() {
	wg := &sync.WaitGroup{}
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for v := range ch1 {
			ch2 <- v * 2
		}
	}()

	go func() {
		for v := range ch2 {
			wg.Done()
			fmt.Printf("x * 2 = %v\n", v)
		}
	}()

	for _, v := range nums {
		wg.Add(1)
		ch1 <- v
	}

	wg.Wait()
	close(ch1)
	close(ch2)
	fmt.Println("All operations completed")
}
