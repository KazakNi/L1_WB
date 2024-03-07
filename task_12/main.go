package main

import "fmt"

/* Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество. */

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	res := getMapSet(arr)
	fmt.Println(res)
}

//создаем сет из мапы с булевым значением
func getMapSet(arr []string) map[string]bool {
	set := make(map[string]bool)
	for _, val := range arr {
		if _, ok := set[val]; !ok {
			set[val] = true
		}
	}
	return set
}
