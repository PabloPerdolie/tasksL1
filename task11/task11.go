package main

import (
	"fmt"
)

// Реализовать пересечение двух неупорядоченных множеств.

func main() {
	array1 := []int{3, 4, 6, 9}    // Первое множество
	array2 := []int{2, 1, 3, 6, 9} // Второе мноджество
	cur := make(map[int]struct{})  // Мапа для записи первого множества
	res := []int{}                 // Результирующий массив

	// Цикл для записи первого массива в мапу
	for _, val := range array1 {
		cur[val] = struct{}{}
	}
	// Цикл для записи в результирующий массив одинаковых элементов
	for _, val := range array2 {
		_, ok := cur[val]
		if ok {
			res = append(res, val)
		}
	}

	// Вывод в консоль
	fmt.Printf("set1: %v", array1)
	fmt.Printf("\nset2: %v", array2)
	fmt.Printf("\nresult: %v", res)
}
