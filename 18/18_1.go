package main

import (
	"fmt"
	"sync"
)

/*
	№ 18 (1-ое решение)

	Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
	По завершению программа должна выводить итоговое значение счетчика.
*/

func main() {
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}

	counter := 0

	for index := 0; index < 1000; index++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			mu.Lock()
			defer mu.Unlock()

			counter++
		}()
	}

	wg.Wait()

	fmt.Println(counter)
}
