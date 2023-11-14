package main

import (
	"fmt"
	"reflect"
)

/*
	№ 14 (2-ое решение)

	Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.
*/

func PrintTypeOf(value any) {
	typeof := reflect.TypeOf(value)

	switch value.(type) {
	case int:
		fmt.Println(typeof)
	case string:
		fmt.Println(typeof)
	case bool:
		fmt.Println(typeof)
	case chan int:
		fmt.Println(typeof)
	case chan string:
		fmt.Println(typeof)
	default:
		fmt.Println("Не удалось определить конкретный тип.")
	}
}

func main() {
	channel := make(chan int)

	PrintTypeOf(42)
	PrintTypeOf("Hello, World!")
	PrintTypeOf(true)
	PrintTypeOf(channel)
}
