package main

import "fmt"

/* Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений. */

func main() {
	fmt.Println(sumOfSquare([]int{2, 4, 6, 8, 10}))
}

func sumOfSquare(array []int) int {
	res := make(chan int)
	done := make(chan bool)
	ans := 0
	go func() {
		defer close(res)
		for _, val := range array {
			res <- val * val

		}
	}()
	go func() {
		for v := range res {
			ans += v
		}
		done <- true
	}()
	<-done
	return ans
}
