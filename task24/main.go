package main

import (
	"fmt"
	"math"
)

// в рамках одного пакета мы можем работать с приватными полями структуры, но в другом
// пакете мы не сможем обратиться к полям х и у

type Point struct {
	x int
	y int
}

func NewPoint(x, y int) *Point {
	return &Point{x: x, y: y}
}

func CalculateDistance(p1, p2 *Point) float64 {
	return math.Sqrt(float64(p1.x-p2.x)*float64(p1.x-p2.x) + float64(p1.y-p2.y)*float64(p1.y-p2.y))
}

func main() {
	point1 := NewPoint(0, 0)
	point2 := NewPoint(4, 3)

	fmt.Println("Calculated distance: ", CalculateDistance(point1, point2))
}
