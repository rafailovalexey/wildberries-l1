package main

import (
	"fmt"
)

/*
	№ 14 (2-ое решение)

	Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.
*/

func PrintTypeOf(value any) {
	switch value.(type) {
	case int:
		fmt.Printf("int\n")
	case string:
		fmt.Printf("string\n")
	case bool:
		fmt.Printf("bool\n")
	case chan int:
		fmt.Printf("chan int\n")
	case chan string:
		fmt.Printf("chan string\n")
	default:
		fmt.Printf("could not determine specific type")
	}
}

func main() {
	channel := make(chan int)

	PrintTypeOf(42)
	PrintTypeOf("Hello, World!")
	PrintTypeOf(true)
	PrintTypeOf(channel)
}
