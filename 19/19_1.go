package main

import "fmt"

/*
	№ 19 (1 решение)

	Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.
*/

func main() {
	value := "главрыба — абырвалг"

	reversed := reverse(value)

	fmt.Printf("source string: %s\n", value)
	fmt.Printf("reversed string: %s\n", reversed)
}

func reverse(value string) string {
	runes := []rune(value)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
