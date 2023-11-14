package main

import (
	"fmt"
	"sort"
)

/*
	№ 17 (1-ое решение)

	Реализовать бинарный поиск встроенными методами языка.
*/

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	target := 6

	index := sort.Search(len(array), func(index int) bool {
		return array[index] >= target
	})

	if index < len(array) && array[index] == target {
		fmt.Println(array[index])
	} else {
		fmt.Println("element not found")
	}
}
