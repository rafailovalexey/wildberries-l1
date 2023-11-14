package main

import "fmt"

/*
	№ 10 (1-ое решение)

	Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
	Объединить данные значения в группы с шагом в 10 градусов.
	Последовательность в подмножноствах не важна.
	Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/

func main() {
	array := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	dictionary := make(map[int][]float32)

	for _, value := range array {
		group := int(value) - (int(value) % 10)

		dictionary[group] = append(dictionary[group], value)
	}

	fmt.Println(dictionary)
}
