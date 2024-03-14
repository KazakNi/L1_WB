package main

import "fmt"

/* Удалить i-ый элемент из слайса. */

func main() {
	a := []string{"A", "B", "C", "D", "E"}
	i := 2

	// Remove the element at index i from a.
	a[i] = a[len(a)-1] // Copy last element to index i.
	a[len(a)-1] = ""   // Erase last element (write zero value).
	a = a[:len(a)-1]   // Truncate slice.

	fmt.Println(a) // [A B E D]

	b := []string{"A", "B", "C", "D", "E"}
	j := 2

	// Remove the element at index i from a.
	copy(b[j:], b[j+1:]) // Shift a[i+1:] left one index.
	b[len(b)-1] = ""     // Erase last element (write zero value).
	b = b[:len(b)-1]     // Truncate slice.

	fmt.Println(b) // [A B D E]
}
