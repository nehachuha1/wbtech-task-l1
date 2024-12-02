package main

import "fmt"

// Функция регистронезависимая (т.е. для неё "А" и "а" - одинаковые символы). Это достигается сдвигом по таблице ASCII
// на +32, если это большие латинские буквы
func CheckUniqueSymbols(input string) bool {
	symbols := make(map[int32]struct{})
	for _, v := range input {
		if v >= 65 && v <= 90 {
			v += 32
			if _, isExists := symbols[v]; isExists {
				return false
			} else {
				symbols[v] = struct{}{}
			}
		} else {
			if _, isExists := symbols[v]; isExists {
				return false
			} else {
				symbols[v] = struct{}{}
			}
		}
	}
	return true
}

func main() {
	var input string
	fmt.Println("Введите строку:")
	fmt.Scanf("%s", &input)

	fmt.Printf("Результат: %v", CheckUniqueSymbols(input))
}
