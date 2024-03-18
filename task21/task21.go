package main

import "fmt"

// Реализовать паттерн «адаптер» на любом примере.

// OldGraphicsLibrary представляет структуру старой библиотеки графики
type OldGraphicsLibrary struct{}

// DrawRectangle представляет метод отрисовки прямоугольника в старой библиотеке
func (ogl *OldGraphicsLibrary) DrawRectangle(x1, y1, x2, y2 int) {
	fmt.Println("Old Graphics Library: Drawing rectangle with coordinates:", x1, y1, x2, y2)
}

// NewGraphicsLibrary представляет структуру новой библиотеки графики
type NewGraphicsLibrary struct{}

// DrawBox представляет метод отрисовки прямоугольника в новой библиотеке
func (ngl *NewGraphicsLibrary) DrawBox(x, y, width, height int) {
	fmt.Println("New Graphics Library: Drawing box with coordinates:", x, y, "and dimensions:", width, height)
}

// Adapter представляет адаптер для использования новой библиотеки вместо старой
type Adapter struct {
	newLibrary *NewGraphicsLibrary
}

// DrawRectangle представляет метод адаптера для отрисовки прямоугольника с использованием новой библиотеки
func (a *Adapter) DrawRectangle(x1, y1, x2, y2 int) {
	width := x2 - x1
	height := y2 - y1
	a.newLibrary.DrawBox(x1, y1, width, height)
}

func main() {
	// Создаем экземпляр адаптера, использующего новую библиотеку
	adapter := &Adapter{
		newLibrary: &NewGraphicsLibrary{},
	}

	// Используем адаптер для отрисовки прямоугольника
	adapter.DrawRectangle(10, 10, 50, 30)
}
