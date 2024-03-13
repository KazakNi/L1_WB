package main

import (
	"fmt"
	"math/rand"
)

/* Реализовать быструю сортировку массива (quicksort) встроенными методами языка. */

func partition(arr []int, l, r int) int {
	i, j := l, r
	pivot := arr[rand.Intn(r)]
	for i <= j {
		for arr[i] < pivot {
			i += 1
		}
		for arr[j] > pivot {
			j -= 1
		}
		if i >= j {
			return i
		}

		arr[i], arr[j] = arr[j], arr[i]
		// edge case for duplicating elements
		if arr[j] == pivot {
			i++
		} else {

			if arr[i] == pivot {
				j--
			}
		}

	}
	return i

}

func quicksort(arr []int, l, r int) {
	if l < r {
		p := partition(arr, l, r)
		quicksort(arr, l, p)
		quicksort(arr, p+1, r)
	}

}

func main() {

	arr := []int{1, 0, 5, 2, 10}
	arr1 := []int{11, -10, 125, 2, -111}
	arr2 := []int{1, 1, 1, 1, -10}
	quicksort(arr, 0, len(arr)-1)
	fmt.Println(arr)
	quicksort(arr1, 0, len(arr1)-1)
	fmt.Println(arr1)
	quicksort(arr2, 0, len(arr2)-1)
	fmt.Println(arr2)

}
