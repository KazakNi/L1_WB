package main

import (
	"fmt"
	"reflect"
)

/* Разработать программу, которая в рантайме способна определить тип переменной:
int, string, bool, channel из переменной типа interface{}. */

func main() {
	v1 := 1
	v2 := "test"
	v3 := []int{1, 3, 5}
	v4 := false

	fmt.Println(defineType(v1))
	fmt.Println(defineType(v2))
	fmt.Println(defineType(v3))
	fmt.Println(defineType(v4))

}

func defineType(type_ interface{}) interface{} {
	return reflect.TypeOf(type_).String()
}
