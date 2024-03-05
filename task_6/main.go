package main

import (
	"context"
	"fmt"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/

func main() {
	stop := make(chan bool)
	// 1 variant with sending value or closing channel
	go goSomeJob(stop)
	<-time.After(time.Second * 1)
	// or close(stop)
	stop <- true
	<-time.After(time.Second * 2)

	// 2 variant with Context package

	ctx, cancel := context.WithCancel(context.Background())
	go goSomeJobContext(ctx)
	<-time.After(time.Second * 1)
	cancel()
	<-time.After(time.Second * 1)
}

func goSomeJob(stop chan bool) {
	for {
		select {
		case <-stop:
			fmt.Println("Stop signal, exiting goroutine 1")
			return
		default:
			fmt.Println("Working goroutine 1")
			time.Sleep(time.Second * 2)
		}
	}
}

func goSomeJobContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Ctx is canceled, exiting goroutine 2")
			return
		default:
			fmt.Println("Doing job goroutine 2")
			time.Sleep(time.Millisecond * 1500)
		}
	}
}
