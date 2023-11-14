package main

import (
	"fmt"
	"sort"
)

/*
	№ 11 (2-ое решение)

	Реализовать пересечение двух неупорядоченных множеств.
*/

func main() {
	array1 := []int{1, 3, 2, 4}
	array2 := []int{5, 4, 3, 6}

	result := intersection(array1, array2)

	fmt.Println("Пересечение множеств в отсортированных массивах:", result)
}

func intersection(array1 []int, array2 []int) []int {
	result := make([]int, 0, 10)

	sort.Ints(array1)
	sort.Ints(array2)

	i, j := 0, 0

	for i < len(array1) && j < len(array2) {
		if array1[i] == array2[j] {
			result = append(result, array1[i])

			i++
			j++
		} else if array1[i] < array2[j] {
			i++
		} else if array1[i] > array2[j] {
			j++
		}
	}

	return result
}
