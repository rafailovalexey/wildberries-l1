package main

import (
	"fmt"
	"math"
	"sync"
)

/*
	№ 3 (1-ое решение)

	Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов с использованием конкурентных вычислений.
*/

func main() {
	mu := &sync.Mutex{}
	wg := &sync.WaitGroup{}

	sequence := []int{2, 4, 6, 8, 10}

	result := 0

	for _, value := range sequence {
		wg.Add(1)

		go func(value int) {
			defer wg.Done()

			square := int(math.Pow(float64(value), 2))

			mu.Lock()
			defer mu.Unlock()

			result += square
		}(value)
	}

	wg.Wait()

	fmt.Println(result)
}
