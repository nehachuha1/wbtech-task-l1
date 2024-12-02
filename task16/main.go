package main

import "fmt"

func partition(numbers []int, low, high int) ([]int, int) {
	pivot := numbers[high]
	i := low
	for j := low; j < high; j++ {
		if numbers[j] < pivot {
			numbers[i], numbers[j] = numbers[j], numbers[i]
			i++
		}
	}
	numbers[i], numbers[high] = numbers[high], numbers[i]
	return numbers, i
}

func QuickSort(numbers []int, low, high int) []int {
	if low < high {
		numbers, p := partition(numbers, low, high)
		numbers = QuickSort(numbers, low, p-1)
		numbers = QuickSort(numbers, p+1, high)
	}
	return numbers
}

func main() {
	numbers := []int{1, 9, 4, 0, 2, 7}
	fmt.Println(QuickSort(numbers, 0, len(numbers)-1))
}
