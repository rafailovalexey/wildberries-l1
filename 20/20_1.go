package main

import (
	"fmt"
	"strings"
)

/*
	№ 20 (1-ое решение)

	Разработать программу, которая переворачивает слова в строке.
	Пример: «snow dog sun — sun dog snow»
*/

func main() {
	value := "snow dog sun"

	result := reverse(value)

	fmt.Println(result)
}

func reverse(value string) string {
	array := strings.Split(value, " ")

	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}

	return strings.Join(array, " ")
}
