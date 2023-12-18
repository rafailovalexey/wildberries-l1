package main

import (
	"log"
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

	log.Printf("program has been stopped\n")
}

func worker(channel <-chan struct{}) {
	log.Printf("worker: start\n")

	for {
		select {
		case <-channel:
			log.Printf("worker: stopped\n")
			return
		default:
			log.Printf("worker: running\n")
			time.Sleep(1 * time.Second)
		}
	}
}
