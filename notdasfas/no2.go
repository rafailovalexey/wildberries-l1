package main

import (
	"fmt"
	"math"
)

func main() {
	numbers := make([]int, 0, 10)
	numbers = append(numbers, 2, 4, 6, 8, 10)

	for _, number := range numbers {
		square := math.Pow(float64(number), 2)

		fmt.Println(square)
	}
}
