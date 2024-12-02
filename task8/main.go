package main

import "fmt"

func main() {
	fmt.Println("Enter a num: ")
	var number int
	fmt.Scanf("%d", &number)
	fmt.Println("Enter a bit index: ")
	var bitIndex int
	fmt.Scanf("%d", &bitIndex)
	fmt.Println("Write a new value")
	var bitValue int
	fmt.Scanf("%d", &bitValue)

	if bitValue == 0 {
		// bitwise operator AND NOT
		//1 &^ 1 = 0
		//1 &^ 0 = 1
		//0 &^ 1 = 0
		//0 &^ 0 = 0
		number = number &^ (1 << bitIndex)
	} else {
		number = number | (1 << bitIndex)
	}
}
