package main

import (
	"fmt"
	"strings"
)

/*
	№ 15 (1 решение)

	К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

	var justString string
	func someFunc() {
	  v := createHugeString(1 << 10)
	  justString = v[:100]
	}

	func main() {
	  someFunc()
	}
*/

var justString string

func someFunc() {
	v := createHugeString(1 << 10)

	justString = v[:100]

	fmt.Println(len(v))
	fmt.Println(len(justString))
}

func createHugeString(size int) string {
	return strings.Repeat("a", size)
}

func main() {
	someFunc()
}
