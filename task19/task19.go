package main

import "fmt"

// Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
// Символы могут быть unicode.

func main() {
	s := "главрыба" // Исходная строка
	var res string  // Результирующая строка
	// Обратный цикл для записи результирующей строки
	for i := len(s) - 1; i >= 0; i-- {
		res += string(s[i])
	}
	// Вывод в консоль
	fmt.Println(res)
}
