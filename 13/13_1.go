package main

import "log"

/*
	№ 13 (1-ое решение)

	Поменять местами два числа без создания временной переменной.
*/

func main() {
	a := 1
	b := 2

	a, b = b, a

	log.Printf("%d\n", a)
	log.Printf("%d\n", b)
}
