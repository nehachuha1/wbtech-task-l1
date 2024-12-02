package main

import (
	"fmt"
	"time"
)

func myOwnSleep(sec int) {
	start := time.Now()

	for {
		if time.Since(start).Seconds() > float64(sec) {
			fmt.Println("Exited from sleeper")
			return
		}
	}
}

func main() {
	sleepTime := 10
	fmt.Println("Start sleep...")
	myOwnSleep(sleepTime)
	fmt.Println("Timer has ringed")
}
