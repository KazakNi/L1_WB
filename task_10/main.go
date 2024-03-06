package main

import "fmt"

/* Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна. */

func main() {
	mapp := make(map[int][]float32)
	var arr = [8]float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	for _, val := range arr {
		mapp[int(val/10)*10] = append(mapp[int(val/10)*10], float32(val))
	}
	for key, val := range mapp {

		fmt.Printf("%v: %v; ", key, val)
	}
}
