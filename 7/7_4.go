package main

import (
	"fmt"
	"sync"
)

/*
	№ 7 (4-ое решение)

	Реализовать конкурентную запись данных в map.
*/

func main() {
	dictionary := make(map[int]struct{})

	wg := &sync.WaitGroup{}
	channelAsMutex := make(chan struct{}, 1)

	for index := 0; index < 1000; index++ {
		wg.Add(1)

		go func(index int) {
			defer wg.Done()

			channelAsMutex <- struct{}{}
			dictionary[index] = struct{}{}

			_ = <-channelAsMutex
		}(index)
	}

	wg.Wait()

	fmt.Printf("%v\n", dictionary)
}
