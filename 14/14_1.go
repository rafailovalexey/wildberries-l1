package main

import (
	"fmt"
	"reflect"
)

/*
	№ 14 (1 решение)

	Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.
*/

func main() {
	var integer any = 1
	var str any = "string"
	var boolean any = true
	var channel any = make(chan string)

	fmt.Println(reflect.TypeOf(integer))
	fmt.Println(reflect.TypeOf(str))
	fmt.Println(reflect.TypeOf(boolean))
	fmt.Println(reflect.TypeOf(channel))
}
