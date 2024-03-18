package main

import "fmt"

// Поменять местами два числа без создания временной переменной.

func main() {
	a := 12
	b := 15
	fmt.Printf("a = %d, b = %d", a, b)
	a, b = b, a
	fmt.Printf("\na = %d, b = %d", a, b)
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("\na = %d, b = %d", a, b)
}
