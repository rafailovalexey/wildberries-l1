package main

import (
	"fmt"
	"time"
)

/*
	№ 25 (3-е решение)

	Реализовать собственную функцию sleep.
*/

func main() {
	seconds := 5

	fmt.Printf("program started\n")
	fmt.Printf("waiting %d seconds\n", seconds)

	sleep(seconds)

	fmt.Printf("program terminated\n")
}

func sleep(seconds int) {
	channel := make(chan struct{})

	go func() {
		time.Sleep(time.Duration(seconds) * time.Second)

		channel <- struct{}{}
	}()

	<-channel
}
