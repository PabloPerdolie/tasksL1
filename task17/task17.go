package main

import "fmt"

// Реализовать бинарный поиск встроенными методами языка.

func binarySearch(arr []int, x int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == x {
			return mid
		} else if arr[mid] < x {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	arr := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20} // Отсортированный массив
	x := 20                                          // Элемент для поиска
	// Выполняем бинарный поиск
	index := binarySearch(arr, x)
	if index != -1 {
		fmt.Printf("Элемент %d найден в позиции %d\n", x, index)
	} else {
		fmt.Printf("Элемент %d не найден в массиве\n", x)
	}
}
