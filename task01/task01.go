package main

import "fmt"

//Задача №1
//Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

// Human - структура с полями Name, Weight, Height, Age
type Human struct {
	Name   string
	Weight int
	Height int
	Age    int
}

// Description - метод структуры Human, выводящий свойства структуры
func (h *Human) Description() {
	fmt.Printf("My name is %s. I'm %d years old. My weight and height are %d and %d", h.Name, h.Age,
		h.Weight, h.Height)
}

// Action - структура, в которую встраивается Human
type Action struct {
	Human
}

func main() {

	// Инициализация переменной action
	action := Action{Human: Human{
		Name:   "Pablo",
		Weight: 75,
		Height: 190,
		Age:    21,
	}}

	// Вызов метода Description структуры Human от структуры Action
	action.Description()
}
