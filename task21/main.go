package main

import "fmt"

// Есть старый рисовальщик, который имеет методы для рисования определенных фигур
type OldDrawer struct{}

func (od *OldDrawer) DrawSquare(x, y, a int) {
	fmt.Printf("Рисую квадрат на координатам X %v и Y %v со стороной а %v", x, y, a)
}

// А есть новый рисовальщик, который работает с интерфейсом Shape. И каждая фигура знает, как себя нарисовать
type Shape interface {
	Draw()
}

// А теперь делаем адаптер
type SquareAdapter struct {
	OldDrawer
	X, Y, A int
}

func (sq *SquareAdapter) Draw() {
	sq.DrawSquare(sq.X, sq.Y, sq.A)
}

func main() {
	square := &SquareAdapter{
		OldDrawer: OldDrawer{},
		X:         5,
		Y:         5,
		A:         10,
	}

	// А сюда мы будем класть разные фигуры для нового рисовальщика. Пока что положим только наш квадрат, который
	// на самом деле из старого рисовальщика, но покрыт адаптером к новому
	newShapes := []Shape{square}

	for _, shape := range newShapes {
		shape.Draw() // в этом цикле используются фигуры нового рисовальщика, в том числе и наша, покрытая адаптером
	}
}
