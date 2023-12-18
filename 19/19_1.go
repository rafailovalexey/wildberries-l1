package main

import "log"

/*
	№ 19 (1-ое решение)

	Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.
*/

func main() {
	value := "главрыба — абырвалг"

	reversed := reverse(value)

	log.Printf("source string: %s\n", value)
	log.Printf("reversed string: %s\n", reversed)
}

func reverse(value string) string {
	runes := []rune(value)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
