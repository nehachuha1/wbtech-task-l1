package main

import "fmt"

func FindIntersection(list1 []int, list2 []int) []int {
	elemsInFirst := make(map[int]struct{})

	for _, v := range list1 {
		if _, isExists := elemsInFirst[v]; !isExists {
			elemsInFirst[v] = struct{}{}
		}
	}

	resultSlice := make([]int, 0)
	for _, v := range list2 {
		if _, isExists := elemsInFirst[v]; isExists {
			resultSlice = append(resultSlice, v)
		}
	}
	return resultSlice
}

func main() {
	lst1 := []int{1, 2, 3, 4, 5}
	lst2 := []int{2, 3, 4, 7, 8}
	fmt.Printf("%#v\n", FindIntersection(lst1, lst2))
}
