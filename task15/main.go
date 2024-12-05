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

// изначально проблема была в том, что мы использовали срез из 100 символов для глобальной переменной для нашего
// функционала, но при этом проинициализированная огромная строка оставалась в памяти независимо от того,
// нужна или нет нам была бы её оставшаяся часть, которую мы не взяли срез
// таким образом, при создании нового слайса мы копируем значения элементов строки в наш слайс байтов
// и далее с ним работаем
func main() {
	someFunc()
	fmt.Println(justString)
}
