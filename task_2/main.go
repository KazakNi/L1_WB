package main

import (
	"fmt"
	"sync"
)

/* Написать программу, которая конкурентно рассчитает значение квадратов
   чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout. */

func main() {
	squareCalc([]int{2, 4, 6, 8, 10})
}

func squareCalc(array []int) {
	var wg sync.WaitGroup
	for _, val := range array {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			fmt.Printf("%v ", val*val)
		}(val)
	}
	wg.Wait()
}
