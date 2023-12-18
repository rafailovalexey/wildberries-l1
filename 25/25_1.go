package main

import (
	"log"
	"time"
)

/*
	№ 25 (1-ое решение)

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
	select {
	case <-time.After(time.Duration(seconds) * time.Second):
		return
	}
}
