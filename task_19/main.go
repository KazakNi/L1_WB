package main

import "fmt"

/* Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode. */

func main() {
	s := "главрыба"
	s1 := "nice гайз"
	res := reverseString(s)
	res2 := reverseString(s1)
	fmt.Println(res, res2)
}

func reverseString(input string) string {
	res := []rune(input)
	i, j := 0, len(res)-1
	for i <= j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return string(res)
}
