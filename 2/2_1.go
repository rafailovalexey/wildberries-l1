package main

import (
	"log"
	"math"
	"sync"
)

/*
	№ 2 (1-ое решение)

	Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	wg := &sync.WaitGroup{}

	for _, number := range numbers {
		wg.Add(1)

		go func(number int) {
			defer wg.Done()

			square := int(math.Pow(float64(number), 2))

			log.Printf("square of the number %d: %d\n", number, square)
		}(number)
	}

	wg.Wait()
}
