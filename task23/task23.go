package main

import "fmt"

// Удалить i-ый элемент из слайса.

func main() {
	arr := []int{1, 5, 3, 6, 8, 2, 4}
	fmt.Println(arr)
	var i int
	fmt.Scan(&i)
	arr = func(array []int, i int) []int {
		return append(array[:i], array[i+1:]...)
	}(arr, i)
	fmt.Println(arr)
}
