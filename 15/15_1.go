package main

import (
	"fmt"
	"strings"
)

/*
	№ 15 (1-ое решение)

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

/*
	Garbage collector может очистить память переменной v после выхода из функции someFunc.
	justString ссылается на 100 первых элементов строки v, строки реализованы через slice,
	соответственно изменение массива или затирание его в памяти приведёт к неправильному/непредсказуемому значению в переменной justString
*/

var justString string

func someFunc() {
	v := createHugeString(1 << 10)

	justString = strings.Clone(v[:100])
}

func createHugeString(size int) string {
	return strings.Repeat("a", size)
}

func main() {
	someFunc()

	fmt.Println(justString)
}
