package main

import "fmt"

/* Реализовать бинарный поиск встроенными методами языка. */

func main() {
	arr := []int{-20, -10, 1, 2, 9, 12, 100, 300, 400}
	target := 9
	fmt.Print(binarySearch(arr, target))
}

func binarySearch(arr []int, target int) (idx int) {
	l, r := 0, len(arr)-1

	for l <= r {
		med := (l + r) / 2
		if arr[med] != target {
			if arr[med] > target {
				r = med - 1
			} else {
				l = med + 1
			}
		} else {
			return med
		}
	}
	return -1
}
