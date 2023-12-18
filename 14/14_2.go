package main

import (
	"log"
)

/*
	№ 14 (2-ое решение)

	Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.
*/

func PrintTypeOf(value any) {
	switch value.(type) {
	case int:
		log.Printf("int\n")
	case string:
		log.Printf("string\n")
	case bool:
		log.Printf("bool\n")
	case chan int:
		log.Printf("chan int\n")
	case chan string:
		log.Printf("chan string\n")
	default:
		log.Printf("could not determine specific type")
	}
}

func main() {
	channel := make(chan int)

	PrintTypeOf(42)
	PrintTypeOf("Hello, World!")
	PrintTypeOf(true)
	PrintTypeOf(channel)
}
