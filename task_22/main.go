package main

import (
	"fmt"
	"math/big"
)

/* Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.
 */
type LongArithmetic struct {
	x string
	y string
}

func (l *LongArithmetic) Multiply() *big.Int {
	x, _ := big.NewInt(0).SetString(l.x, 10)
	y, _ := big.NewInt(0).SetString(l.y, 10)
	return big.NewInt(0).Mul(x, y)
}

func (l *LongArithmetic) Division() *big.Int {
	x, _ := big.NewInt(0).SetString(l.x, 10)
	y, _ := big.NewInt(0).SetString(l.y, 10)
	return big.NewInt(0).Div(x, y)
}

func (l *LongArithmetic) Add() *big.Int {
	x, _ := big.NewInt(0).SetString(l.x, 10)
	y, _ := big.NewInt(0).SetString(l.y, 10)
	return big.NewInt(0).Add(x, y)
}

func (l *LongArithmetic) Subtraction() *big.Int {
	x, _ := big.NewInt(0).SetString(l.x, 10)
	y, _ := big.NewInt(0).SetString(l.y, 10)
	return big.NewInt(0).Sub(x, y)
}

func main() {
	l := LongArithmetic{x: "10000000000000000000000000", y: "20000000000000000000000000000000000000000000000"}
	fmt.Println(l.Multiply())

}
