package main

import (
	"fmt"
	"math"
)

/*
	№ 2 (2-ое решение)

	Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	channel := make(chan string)

	go func() {
		for _, number := range numbers {
			square := int(math.Pow(float64(number), 2))

			channel <- fmt.Sprintf("Квадрат числа %d: %d\n", number, int(square))
		}

		close(channel)
	}()

	for value := range channel {
		fmt.Print(value)
	}
}
