package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
	№ 26 (1 решение)

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
	dictionary := make(map[string]struct{})

	value = strings.ToLower(value)

	for _, char := range value {
		temporary := strconv.Itoa(int(char))

		if _, isExist := dictionary[temporary]; isExist {
			return false
		}

		dictionary[temporary] = struct{}{}
	}

	return true
}
