package main

import (
	"log"
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
		defer close(channel)

		for _, number := range numbers {
			square := int(math.Pow(float64(number), 2))

			channel <- log.Sprintf("square of the number %d: %d\n", number, square)
		}
	}()

	for value := range channel {
		log.Printf("%s\n", value)
	}
}
