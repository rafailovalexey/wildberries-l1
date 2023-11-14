package main

import "fmt"

/*
	№ 13 (1 решение)

	Поменять местами два числа без создания временной переменной.
*/

func main() {
	a := 1
	b := 2

	a, b = b, a

	fmt.Println(a)
	fmt.Println(b)
}
