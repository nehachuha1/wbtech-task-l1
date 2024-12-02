package main

import "fmt"

func main() {
	var i int
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fmt.Println("Enter an index that should be deleted")
	fmt.Scanf("%d", &i)
	if i >= len(numbers) {
		fmt.Println("You wrote index that bigger than length of numbers")
	} else {
		fmt.Printf("Deleted element: %v\n", numbers[i])
		numbers = append(numbers[:i], numbers[i+1:]...)
	}
	fmt.Printf("Result: %#v\n", numbers)
}
