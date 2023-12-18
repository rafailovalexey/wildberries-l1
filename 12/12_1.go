package main

import "log"

/*
	№ 12 (1-ое решение)

	Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/

func main() {
	array := []string{"cat", "cat", "dog", "cat", "tree"}

	result := intersaction(array)

	log.Printf("intersection of many %v\n", result)
}

func intersaction(array []string) map[any]any {
	dictionary := make(map[string]struct{}) // Является собственным множеством
	result := make(map[any]any)

	for _, element := range array {
		if _, isExist := dictionary[element]; isExist {
			result[element] = struct{}{}
		}

		if _, isExist := dictionary[element]; !isExist {
			dictionary[element] = struct{}{}
		}
	}

	return result
}
