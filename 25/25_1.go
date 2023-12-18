package main

import (
	"fmt"
	"time"
)

/*
	№ 25 (1-ое решение)

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
	select {
	case <-time.After(time.Duration(seconds) * time.Second):
		return
	}
}
