package main

import (
	"fmt"
	"math"
)

/* Разработать программу нахождения расстояния между двумя точками,
   которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором. */

type Point struct {
	x float64
	y float64
}

func (p *Point) Distance(pp Point) int {
	return int(math.Sqrt((p.x-pp.x)*(p.x-pp.x) + (p.y-pp.y)*(p.y-pp.y)))
}

func main() {
	p1 := Point{x: 1, y: 2}
	p2 := Point{x: 10, y: 10}
	fmt.Println(p1.Distance(p2))
}
