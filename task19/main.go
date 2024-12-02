package main

import "fmt"

func ReverseString(input string) string {
	convertedString := []rune(input)
	for i := 0; i < len(convertedString)/2; i++ {
		convertedString[i], convertedString[len(convertedString)-i-1] = convertedString[len(convertedString)-i-1],
			convertedString[i]
	}
	return string(convertedString)
}

func main() {
	input := "masndnokas#😎😎😎😎😎😎"
	fmt.Printf("Reversed string: %v", ReverseString(input))
}
