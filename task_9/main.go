package main

import (
	"fmt"
	"sync"
)

/* Разработать конвейер чисел. Даны два канала:
в первый пишутся числа (x) из массива, во второй — результат операции x*2,
после чего данные из второго канала должны выводиться в stdout. */

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	sample := []int{1, 2, 3, 4, 5}
	getNums := make(chan int, 1)
	multiRes := make(chan int, 1)

	go loadNums(getNums, sample, &wg)
	go multiplyNums(getNums, multiRes, &wg)

	wg.Wait()

}

func loadNums(src chan int, arr []int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range arr {
		src <- val
	}
	close(src)
}

func multiplyNums(src chan int, res chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range src {
		res <- val * 2
		fmt.Print(" ", <-res)
	}
	close(res)
}
