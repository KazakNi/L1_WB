package main

import (
	"fmt"
	"sync/atomic"
)

/* Поменять местами два числа без создания временной переменной.
 */

func main() {
	// first variant
	var a atomic.Int32
	a.Store(1)
	var b int32 = 2
	b = a.Swap(b)
	fmt.Printf("a = %v, b = %v\n", a.Load(), b)

	// second variant

	c, d := 2, 3
	c, d = d, c
	fmt.Printf("c = %v, d = %v\n", c, d)

	// third variant behind the second variant

	c = d ^ c
	d = c ^ d
	c = d ^ c
	fmt.Printf("c = %v, d = %v", c, d)
}
