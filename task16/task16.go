package main

import "fmt"

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

func quickSort(arr []int, first, last int) {
	if len(arr) <= 1 {
		return
	}
	i := first
	j := last
	x := arr[(i+j)/2]
	for {
		for i < last && arr[i] < x {
			i++
		}
		for j > first && arr[j] > x {
			j--
		}
		if i <= j {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
			i++
			j--
		} else {
			break
		}
	}
	if i < last {
		quickSort(arr, i, last)
	}
	if j > first {
		quickSort(arr, first, j)
	}
}

func main() {
	arr := []int{9, 7, 5, 11, 12, 2, 14, 3, 10, 6}
	fmt.Println("Исходный массив:", arr)
	quickSort(arr, 0, len(arr)-1)
	fmt.Println("Отсортированный массив:", arr)
}
