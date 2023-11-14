package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
	№ 18 (2-ое решение)

	Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
	По завершению программа должна выводить итоговое значение счетчика.
*/

func main() {
	wg := &sync.WaitGroup{}

	counter := int64(0)

	for index := 0; index < 1000; index++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			atomic.AddInt64(&counter, 1)
		}()
	}

	wg.Wait()

	fmt.Println(counter)
}
