package main

import (
	"fmt"
	"strings"
)

/* Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
	aabcd — false */

func main() {
	s := "abcd"
	s1 := "abCdefAaf"
	s2 := "aabcd"
	fmt.Println(checkUniqueness(s))
	fmt.Println(checkUniqueness(s1))
	fmt.Println(checkUniqueness(s2))

}

func checkUniqueness(s string) bool {
	m := make(map[string]bool)
	for _, rune_ := range s {
		letter := strings.ToLower(string(rune_))
		if _, ok := m[letter]; !ok {
			m[letter] = true
		} else {
			return false
		}
	}
	return true
}
