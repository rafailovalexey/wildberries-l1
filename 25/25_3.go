package main

import (
	"log"
	"time"
)

/*
	№ 25 (3-е решение)

	Реализовать собственную функцию sleep.
*/

func main() {
	seconds := 5

	log.Printf("program started\n")
	log.Printf("waiting %d seconds\n", seconds)

	sleep(seconds)

	log.Printf("program terminated\n")
}

func sleep(seconds int) {
	channel := make(chan struct{})

	go func() {
		time.Sleep(time.Duration(seconds) * time.Second)

		channel <- struct{}{}
	}()

	<-channel
}
