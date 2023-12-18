package main

import "fmt"

/*
	№ 23 (2 решение)

	Удалить i-ый элемент из слайса.
*/

func main() {
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	array = DeleteArrayByIndex(array, 5)

	fmt.Printf("%v %d %d\n", array, len(array), cap(array))
}

func DeleteArrayByIndex(array []int, index int) []int {
	result := make([]int, 0, len(array))

	for i, value := range array {
		if i == index {
			continue
		}

		result = append(result, value)
	}

	return result
}
