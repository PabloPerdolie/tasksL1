package main

import (
	"fmt"
	"strings"
)

//К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.
//
//var justString string
//func someFunc() {
//  v := createHugeString(1 << 10)
//  justString = v[:100]
//}
//
//func main() {
//  someFunc()
//}
// Ответ: мы ссылаемся на одну и ту же область памяти, что может привести к непреднамеренному изменению данных

var justString string

func createHugeString(n int) string {
	arr := make([]byte, n)
	for i := 0; i < n; i++ {
		arr[i] = byte('A')
	}
	return string(arr)
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = strings.Clone(v[:100])
	fmt.Println(len(justString), justString)
}

func main() {
	someFunc()
}
