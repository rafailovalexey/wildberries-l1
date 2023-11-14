package main

import "fmt"

/*
	№ 11 (1 решение)

	Реализовать пересечение двух неупорядоченных множеств.
*/

func main() {
	map1 := map[any]any{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}}
	map2 := map[any]any{3: struct{}{}, 4: struct{}{}, 5: struct{}{}, 6: struct{}{}}

	result := intersection(map1, map2)

	fmt.Println("Пересечение множеств:", result)
}

func intersection(map1 map[any]any, map2 map[any]any) map[any]any {
	result := make(map[any]any)

	for element := range map1 {
		if _, isExist := map2[element]; isExist {
			result[element] = struct{}{}
		}
	}

	return result
}
