package main

import "fmt"

/*
	№ 18 (4-ое решение)

	Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
	По завершению программа должна выводить итоговое значение счетчика.
*/

func main() {
	channel := make(chan int, 1)

	channel <- 0

	for index := 0; index < 1000; index++ {
		value, isOpen := <-channel

		if !isOpen {
			break
		}

		channel <- value + 1
	}

	close(channel)

	result := <-channel

	fmt.Printf("%d\n", result)
}
