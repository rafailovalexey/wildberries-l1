package main

import (
	"fmt"
	"time"
)

/*
	№ 25 (3 решение)

	Реализовать собственную функцию sleep.
*/

func main() {
	seconds := 5

	fmt.Println("Program started")
	fmt.Printf("Waiting %d seconds\n", seconds)

	sleep(seconds)

	fmt.Println("Program terminated")
}

func sleep(seconds int) {
	channel := make(chan struct{})

	go func() {
		time.Sleep(time.Duration(seconds) * time.Second)

		channel <- struct{}{}
	}()

	<-channel
}
