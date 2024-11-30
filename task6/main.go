package main

import (
	"context"
	"fmt"
	"time"
)

// Ex 1: Остановка горутины через контекст с таймаутом
func GoroutineWithCtxWithTimeout(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("context has been canceled")
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("I'm working right now!")
		}
	}
}

// Ex 2: Остановка горутины через контекст с cancel()
func GoroutineWithCtxWithCancel(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("context has been canceled")
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("I'm working right now!")
		}
	}
}

// Ex3: Остановка горутины, слушающей quit канал
func GoroutineWithChannel(out chan int, quit chan bool) {
	value := 1

	for {
		select {
		case <-quit:
			close(out)
			fmt.Println("Exited from goroutine")
		case out <- value:
			value++
		}
	}
}

func main() {
	// Горутина с контекстом с таймаутом
	fmt.Println("Example 1:")

	ctx := context.Background()
	ctxWithTimeout, _ := context.WithTimeout(ctx, time.Duration(3*time.Second))

	go GoroutineWithCtxWithTimeout(ctxWithTimeout)
	time.Sleep(5 * time.Second)
	fmt.Println("Example 1 completed")

	// Горутина с контекстом без таймаута, но с cancel()
	fmt.Println("Example 2:")

	ctx = context.Background()
	ctxWithCancel, cancel := context.WithCancel(ctx)
	go GoroutineWithCtxWithCancel(ctxWithCancel)

	time.Sleep(5 * time.Second)
	cancel()
	fmt.Println("Example 2 completed")
	time.Sleep(2 * time.Second)

	// Горутина с каналом quit
	fmt.Println("Example 3:")
	out := make(chan int)
	quit := make(chan bool)
	go GoroutineWithChannel(out, quit)

	for v := range out {
		time.Sleep(1 * time.Second)
		fmt.Printf("Current value: %v\n", v)
		if v > 4 {
			quit <- true
			break
		}
	}
	close(quit)
	fmt.Println("Example 3 completed")

	fmt.Println("All examples completed!")
}
