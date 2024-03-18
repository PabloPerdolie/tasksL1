package main

import "fmt"

// Разработать программу, которая в рантайме способна определить тип переменной:
// int, string, bool, channel из переменной типа interface{}.

func Type(inter interface{}) {
	switch v := inter.(type) {
	case int:
		fmt.Println(v, "is int")
	case string:
		fmt.Println(v, "is string")
	case bool:
		fmt.Println(v, "is bool")
	case chan int:
		fmt.Println(v, "is chan int")
	default:
		fmt.Println("idk")
	}
}

func main() {
	var a int
	var b = "vgvhgvhkg"
	var c bool
	d := make(chan int)

	Type(a)
	Type(b)
	Type(c)
	Type(d)
}
