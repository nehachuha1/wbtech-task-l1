package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h *Human) SaySomething() {
	fmt.Println("This is method that related to Human struct")
}

func (h *Human) DoSomething() {
	fmt.Println("Action that related to Human struct")
}

type Action struct {
	Human
	SomeFieldInAction string
}

//func (action *Action) DoSomething() {
//
//}

func main() {
	action := &Action{
		Human:             Human{Name: "Name", Age: 20},
		SomeFieldInAction: "someFiledInAction",
	}
	action.SaySomething()
	// При наличии у структуры Action метода DoSomething вызывался бы сначала он. Так как
	// этот метод реализован только у структуры Human, которая встроена в Action, то вызывается
	//еализация метода DoSomething у структуры Human
	action.DoSomething()
}
