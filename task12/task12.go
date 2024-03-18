package main

import (
	"fmt"
)

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

func main() {
	array := []string{"cat", "cat", "dog", "cat", "tree"} // Исходный массив строк

	m := make(map[string]struct{}) // Промежуточная мапа для исключения повторений

	// Цикл для записи в мапу
	for _, v := range array {
		m[v] = struct{}{}
	}

	res := make([]string, 0, len(m)) // Результирующий массив строк
	// Цикл для записи из мапы в массив
	for k, _ := range m {
		res = append(res, k)
	}

	// Вывод в консоль
	fmt.Printf("set: %v", res)
}
