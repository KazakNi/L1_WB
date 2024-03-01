package main

import (
	"fmt"
	"time"
)

/* Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования). */

type Human struct {
	Name string
	Age  int
	Sex  string
}

func (h Human) Voice() string {
	if h.Sex == "female" {
		return "la-la-la"
	} else {
		return "ho-ho-ho"
	}
}

func (h Human) Run() string {
	if h.Age < 1 {
		return "crawling"
	} else if h.Age > 100 {
		return "walking with stilts"
	}
	return "jogging"
}

type Action struct {
	Human
	Type     string
	Duration time.Duration
}

func (a *Action) PlayFootball(distance *int, fieldLength int) {

	for a.Duration > 0 && *distance <= fieldLength {
		a.KickBall(distance)
		fmt.Printf("Ball was kicked to %v meters by %s!\n", *distance, a.Name)
	}
}

func (a Action) KickBall(distance *int) {

	if a.Age < 2 || a.Age > 70 {
		*distance += 1
	} else {
		*distance += 3
	}

}

func main() {
	playSoccer := Action{
		Human: Human{
			Name: "Andrew",
			Age:  30,
			Sex:  "Male",
		},
		Type:     "football playing",
		Duration: time.Minute * 60,
	}
	distance := 0
	playSoccer.PlayFootball(&distance, 120)
}
