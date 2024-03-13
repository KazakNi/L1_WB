package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

/* Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика. */

func main() {
	var c counter = 0
	for i := 0; i < 1000000; i++ {
		go func() {
			c.inc()
		}()

	}
	<-time.After(time.Second * 1)
	fmt.Println(c.get())

}

type counter int32

func (c *counter) inc() int32 {
	return atomic.AddInt32((*int32)(c), 1)
}

func (c *counter) get() int32 {
	return atomic.LoadInt32((*int32)(c))
}
