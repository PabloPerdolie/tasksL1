package main

import (
	"fmt"
	"math/big"
)

// Разработать программу, которая перемножает, делит, складывает,
// вычитает две числовых переменных a,b, значение которых > 2^20.

func main() {
	a := big.NewFloat(3475637456374563)
	b := big.NewFloat(2342048293482948)
	fmt.Println(new(big.Float).Add(a, b))
	fmt.Println(new(big.Float).Sub(a, b))
	fmt.Println(new(big.Float).Mul(a, b))
	fmt.Println(new(big.Float).Quo(a, b))
}
