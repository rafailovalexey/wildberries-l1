package main

import "fmt"

/*
	№ 13 (1-ое решение)

	Поменять местами два числа без создания временной переменной.
*/

func main() {
	a := 1
	b := 2

	a, b = b, a

	fmt.Printf("%d\n", a)
	fmt.Printf("%d\n", b)
}
