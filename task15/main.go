package main

import "fmt"

func createHugeString(size int) string {
	return string(make([]byte, size))
}

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = string([]byte(v[:100]))
}
func main() {
	someFunc()
	fmt.Println(justString)
}
