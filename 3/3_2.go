package main

import (
	"math"
	"sync"
	"sync/atomic"
)

/*
	№ 3 (2-ое решение)

	Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов с использованием конкурентных вычислений.
*/

func main() {
	wg := &sync.WaitGroup{}

	sequence := []int64{2, 4, 6, 8, 10}

	var result int64 = 0

	for _, value := range sequence {
		wg.Add(1)

		go func(value int64) {
			defer wg.Done()

			square := int64(math.Pow(float64(value), 2))

			atomic.AddInt64(&result, square)
		}(value)
	}

	wg.Wait()

	log.Printf("%s\n", result)
}
