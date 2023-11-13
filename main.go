package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
№ 4 (решение 1)

Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/

func main() {
	workers := runtime.NumCPU()

	channel := make(chan string)
	wg := &sync.WaitGroup{}

	go func() {
		counter := 0

		for {
			Publisher(channel, fmt.Sprintf("Publisher-1: %d", counter))

			counter++
		}
	}()

	for index := 1; index <= workers; index++ {
		wg.Add(1)

		go func(index int) {
			defer wg.Done()

			for message := range channel {
				fmt.Println(fmt.Sprintf("Subscriber-%d: %s", index, message))
			}
		}(index)
	}

	wg.Wait()
}

func Publisher(channel chan string, message string) {
	channel <- message
}
