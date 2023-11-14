package main

import (
	"fmt"
	"time"
)

/*
	№ 6 (4-ое решение)

	Реализовать все возможные способы остановки выполнения горутины.
*/

func main() {
	channel := make(chan struct{})

	go worker(channel)

	time.Sleep(3 * time.Second)
	close(channel)
	time.Sleep(3 * time.Second)

	fmt.Println("Program has been stopped")
}

func worker(channel <-chan struct{}) {
	fmt.Println("Worker: Start")

	for {
		select {
		case <-channel:
			fmt.Println("Worker: Stopped")
			return
		default:
			fmt.Println("Worker: Running")
			time.Sleep(1 * time.Second)
		}
	}
}
