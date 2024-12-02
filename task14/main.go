package main

import (
	"fmt"
	"math/rand"
)

func GetSomeVar() interface{} {
	res := rand.Intn(4)
	switch res {
	case 0:
		return 1
	case 1:
		return "123"
	case 2:
		return true
	default:
		return make(chan interface{})
	}
}

func main() {
	test := GetSomeVar()
	switch test.(type) {
	case int:
		fmt.Println("It's int!")
	case string:
		fmt.Println("It's string!")
	case bool:
		fmt.Println("It's bool!")
	default:
		fmt.Println("It's channel!")
	}
}
