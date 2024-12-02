package main

import "fmt"

func main() {
	a := 10
	b := 5
	fmt.Printf("%v %v\n", a, b)
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("%v %v\n", a, b)
}
