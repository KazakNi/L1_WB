package main

import (
	"fmt"
	"strings"
)

/* Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow». */

func main() {
	s := "snow dog sun"
	fmt.Println(reverseWords(s))

}

func reverseWords(s string) string {
	res := strings.Split(s, " ")
	i, j := 0, len(res)-1

	for i <= j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return strings.Join(res, " ")
}
