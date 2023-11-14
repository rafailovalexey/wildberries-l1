package main

import (
	"fmt"
	"sync"
)

/*
	№ 18 (3 решение)

	Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
	По завершению программа должна выводить итоговое значение счетчика.
*/

func main() {
	wg := &sync.WaitGroup{}
	channelAsMutex := make(chan struct{}, 1)

	counter := 0

	for index := 0; index < 1000; index++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			channelAsMutex <- struct{}{}

			counter++

			_ = <-channelAsMutex
		}()
	}

	wg.Wait()

	fmt.Println(counter)
}
