package main

import (
	"fmt"
	"sort"
)

/*
	№ 16 (1 решение)

	Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

func main() {
	integers := []int{9, 7, 5, 11, 12, 2, 14, 3, 10, 6}
	floats := []float64{3.5, 2.0, 1.7, 4.2, 0.9}
	strs := []string{"apple", "orange", "banana", "grape", "kiwi"}

	sort.Ints(integers)
	sort.Float64s(floats)
	sort.Strings(strs)

	fmt.Println("Отсортированный массив чисел:", integers)
	fmt.Println("Отсортированный массив чисел с плавающей запятой:", floats)
	fmt.Println("Отсортированный массив строк:", strs)
}
