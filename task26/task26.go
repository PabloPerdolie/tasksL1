package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
// Функция проверки должна быть регистронезависимой.
//
// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false

func isUniq(str string) bool {
	str = strings.ToLower(str)
	mp := make(map[rune]bool, len(str))
	for _, v := range str {
		if mp[v] {
			return false
		}
		mp[v] = true
	}
	return true
}

func main() {
	str1 := "abcd"
	fmt.Println(str1, isUniq(str1))
	str2 := "abCdefA"
	fmt.Println(str2, isUniq(str2))
	str3 := "aabcd"
	fmt.Println(str3, isUniq(str3))
}
