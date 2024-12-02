package main

import (
	"fmt"
)

func MakeGroups(temperatures []float64) map[int][]float64 {
	resultMap := make(map[int][]float64)
	for _, v := range temperatures {
		ost := v / 10.0
		if _, isExists := resultMap[int(ost)*10]; !isExists {
			resultMap[int(ost)*10] = []float64{v}
		} else {
			resultMap[int(ost)*10] = append(resultMap[int(ost)*10], v)
		}
	}
	return resultMap
}

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	fmt.Printf("%#v", MakeGroups(temperatures))
}
