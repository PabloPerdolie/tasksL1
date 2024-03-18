package main

import "fmt"

// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

func main() {
	s := "snow dog sun" // Исходная строка
	temp := " "         // Врменная переменная
	var res string
	// Цикл для разбора строки на слова
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			res = temp + res
			temp = " "
		} else {
			temp += string(s[i])
		}
	}
	res = temp + res // Добавление последнего слова

	// Вывод в консоль
	fmt.Println(res[1:])
}
