package main

import "fmt"

func BinarySearch(numbers []int, toFindNumber int) int {
	low := 0
	high := len(numbers) - 1

	answer := -1
	for low <= high {
		mid := low + (high-low)/2
		if numbers[mid] == toFindNumber {
			return mid
		} else if numbers[mid] > toFindNumber {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return answer
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	toFindNum := 4
	fmt.Println(BinarySearch(nums, toFindNum))
	toFindNum = 10
	fmt.Println(BinarySearch(nums, toFindNum))
}
