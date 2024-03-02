package main

import (
	"fmt"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в канал,

	а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
*/
var durationOfScript int

func main() {
	res := make(chan int, 5)

	fmt.Print("Enter time of duration for this programm:")
	fmt.Scanln(&durationOfScript)

	end := time.After(time.Duration(durationOfScript) * time.Second)

	for i := 1; ; i++ {
		select {
		case <-end:
			fmt.Println("Programm ending")
			return
		default:
			res <- i
			fmt.Println("Data from channel: ", <-res)

		}

	}
}
