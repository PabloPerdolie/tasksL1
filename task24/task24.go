package main

import (
	"fmt"
	"math"
)

// Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

type Point struct {
	X, Y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		X: x,
		Y: y,
	}
}

func GetDistance(point1, point2 *Point) float64 {
	dx := point1.X - point2.X
	dy := point1.Y - point2.Y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	point1 := NewPoint(3, 5)
	point2 := NewPoint(4, 6)
	fmt.Println(GetDistance(point1, point2))
}
