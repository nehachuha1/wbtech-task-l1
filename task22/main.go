package main

import "fmt"

func main() {
	var a, b int64
	// в 2^20 * 2^20 содержится всего лишь 13 знаков, а в int64 помещается 18 знаков. Соответственно мы всегда
	// будем помещаться в int64

	fmt.Scanf("%d %d", &a, &b)
	fmt.Printf("a * b = %v\n", a*b)
	fmt.Printf("a / b = %v\n", a/b)
	fmt.Printf("a + b = %v\n", a+b)
	fmt.Printf("a - b = %v\n", a-b)
}
