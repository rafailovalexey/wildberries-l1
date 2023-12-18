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

	fmt.Printf("program has been stopped\n")
}

func worker(channel <-chan struct{}) {
	fmt.Printf("worker: start\n")

	for {
		select {
		case <-channel:
			fmt.Printf("worker: stopped\n")
			return
		default:
			fmt.Printf("worker: running\n")
			time.Sleep(1 * time.Second)
		}
	}
}
