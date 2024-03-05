package main

import (
	"fmt"
)

/* Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0. */

func main() {
	var j int64
	var p int
	fmt.Println("Enter a number")
	fmt.Scanln(&j)
	fmt.Println("Enter a position number for bit switch:")
	fmt.Scanln(&p)
	if p > binaryDigitsCount(j) {
		fmt.Println("Sorry, out of index")
	}
	fmt.Printf("New number is: %v", bitwiseNot(j, p))
}

func bitwiseNot(n int64, pos int) int64 {
	n ^= (1 << pos)
	return n
}

func binaryDigitsCount(n int64) int {
	counter := 0
	for n > 0 {
		n /= 2
		counter++
	}
	return counter
}
