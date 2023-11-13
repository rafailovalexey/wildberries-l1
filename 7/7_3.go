package main

import (
	"fmt"
	"sync"
)

/*
	№ 7 (3 решение)

	Реализовать конкурентную запись данных в map.
*/

func main() {
	dictionary := &sync.Map{}

	wg := &sync.WaitGroup{}

	for index := 0; index < 1000; index++ {
		wg.Add(1)

		go func(index int) {
			defer wg.Done()

			dictionary.Store(index, struct{}{})
		}(index)
	}

	wg.Wait()

	fmt.Println(dictionary)
}
