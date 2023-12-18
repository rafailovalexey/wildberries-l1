package main

import (
	"fmt"
	"sort"
)

/*
	№ 16 (1-ое решение)

	Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

func main() {
	integers := []int{9, 7, 5, 11, 12, 2, 14, 3, 10, 6}
	floats := []float64{3.5, 2.0, 1.7, 4.2, 0.9}
	strs := []string{"apple", "orange", "banana", "grape", "kiwi"}

	sort.Ints(integers)
	sort.Float64s(floats)
	sort.Strings(strs)

	fmt.Printf("sorted array of numbers %v\n", integers)
	fmt.Printf("sorted array of floating point numbers %v\n", floats)
	fmt.Printf("sorted array of strings %v\n", strs)
}
