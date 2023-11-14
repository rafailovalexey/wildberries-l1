package main

import (
	"fmt"
	"strings"
)

/*
	№ 26 (2-ое решение)

	Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

	Например:
	abcd — true
	abCdefAaf — false
	aabcd — false
*/

func main() {
	value := "Aabcd"

	result := CheckCharExists(value)

	fmt.Println(result)
}

func CheckCharExists(value string) bool {
	value = strings.ToLower(value)

	for index, char := range value {
		if char != ' ' && strings.ContainsRune(value[index+1:], char) {
			return false
		}
	}

	return true
}
