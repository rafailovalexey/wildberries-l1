package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
	№ 6 (1-ое решение)

	Реализовать все возможные способы остановки выполнения горутины.
*/

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go worker(ctx)

	go func() {
		exit := make(chan os.Signal)
		signal.Notify(exit, syscall.SIGINT)

		for {
			select {
			case <-ctx.Done():
				return
			case <-exit:
				log.Printf("CTRL+C received. Stopping subscribers...\n")

				cancel()
			}
		}
	}()

	<-ctx.Done()

	log.Printf("program has been stopped\n")
}

func worker(ctx context.Context) {
	log.Printf("worker: start\n")

	for {
		select {
		case <-ctx.Done():
			log.Printf("worker: stopped\n")
			return
		default:
			log.Printf("worker: running\n")
			time.Sleep(1 * time.Second)
		}
	}
}
