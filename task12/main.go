package main

import "fmt"

func GetSet(list1 []string) []string {
	elems := make(map[string]struct{})

	for _, v := range list1 {
		if _, isExists := elems[v]; !isExists {
			elems[v] = struct{}{}
		}
	}

	resultSlice := make([]string, 0)
	for key, _ := range elems {
		resultSlice = append(resultSlice, key)
	}
	return resultSlice
}

func main() {
	lst1 := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Printf("%#v\n", GetSet(lst1))
}
