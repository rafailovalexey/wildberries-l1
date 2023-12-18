package main

import (
	"log"
	"reflect"
)

/*
	№ 14 (1-ое решение)

	Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.
*/

func main() {
	var integer any = 1
	var str any = "string"
	var boolean any = true
	var channel any = make(chan string)

	log.Printf("%s\n", reflect.TypeOf(integer))
	log.Printf("%s\n", reflect.TypeOf(str))
	log.Printf("%s\n", reflect.TypeOf(boolean))
	log.Printf("%s\n", reflect.TypeOf(channel))
}
