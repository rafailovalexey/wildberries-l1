package main

import "log"

/*
	№ 23 (1 решение)

	Удалить i-ый элемент из слайса.
*/

func main() {
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	array = DeleteArrayByIndex(array, 5)

	log.Printf("%v %d %d\n", array, len(array), cap(array))
}

func DeleteArrayByIndex(array []int, index int) []int {
	return append(array[:index], array[index+1:]...)
}
