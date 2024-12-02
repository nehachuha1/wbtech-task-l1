package main

import (
	"fmt"
	"strings"
)

func ReverseWords(input string) string {
	words := strings.Split(input, " ")
	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-1-i] = words[len(words)-1-i], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	input := "snow dog sun"
	fmt.Printf("Input string: %#v\n", input)
	fmt.Printf("Result: %#v\n", ReverseWords(input))
}
