package main

import (
	"log"
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
		log.Printf("%v\n", array[index])
	} else {
		log.Printf("element not found\n")
	}
}
