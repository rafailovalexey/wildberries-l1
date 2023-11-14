package main

import (
	"fmt"
	"time"
)

/*
	№ 25 (2-ое решение)

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
	timer := time.NewTimer(time.Duration(seconds) * time.Second)
	<-timer.C
}
