package main

import (
	"fmt"
	"sync"
)

/*
	№ 7 (1-ое решение)

	Реализовать конкурентную запись данных в map.
*/

func main() {
	dictionary := make(map[int]struct{})

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}

	for index := 0; index < 1000; index++ {
		wg.Add(1)

		go func(index int) {
			defer wg.Done()

			mu.Lock()
			defer mu.Unlock()

			dictionary[index] = struct{}{}
		}(index)
	}

	wg.Wait()

	fmt.Printf("%v\n", dictionary)
}
