package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

/*
Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые

	читают произвольные данные из канала и выводят в stdout.
	Необходима возможность выбора количества воркеров при старте.

	Программа должна завершаться по нажатию Ctrl+C.
	Выбрать и обосновать способ завершения работы всех воркеров.
*/
var workersNum int

func worker(w int, jobs <-chan int) {
	for val := range jobs {
		fmt.Printf("Worker number %v, done: %v\n", w, val)
	}
}

func main() {

	fmt.Println("Enter workers number: ")
	fmt.Scanln(&workersNum)

	termChan := make(chan os.Signal, 1)
	jobs := make(chan int, workersNum)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)

	for w := 1; w <= workersNum; w++ {
		go worker(w, jobs)
	}

	for i := 1; ; i += 10 {
		select {
		case <-termChan:
			close(jobs)
			fmt.Println("Cancellation of programm")
			return
		default:
			jobs <- i

		}
	}

}
