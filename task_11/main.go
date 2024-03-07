package main

import "fmt"

/*
Реализовать пересечение двух неупорядоченных множеств.
*/

func main() {
	arr1 := []string{"cat", "banana", "frog", "pipe"}
	arr2 := []string{"bird", "ginger", "frog", "clone", "pipe"}
	// генерируем множества
	set1 := getMapSet(arr1)
	set2 := getMapSet(arr2)
	// ищем пересечение
	resSet := intersectSets(set1, set2)
	fmt.Println(resSet)

}

func intersectSets(set1, set2 map[string]bool) map[string]bool {
	res := make(map[string]bool)
	for key, _ := range set1 {
		if set2[key] {
			res[key] = true
		}
	}
	return res
}

func getMapSet(arr []string) map[string]bool {
	set := make(map[string]bool)
	for _, val := range arr {
		set[val] = true
	}
	return set
}
