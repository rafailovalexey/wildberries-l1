package main

import (
	"log"
	"time"
)

/*
	№ 25 (2-ое решение)

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
	timer := time.NewTimer(time.Duration(seconds) * time.Second)
	<-timer.C
}
