package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
	№ 6 (3-е решение)

	Реализовать все возможные способы остановки выполнения горутины.
*/

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
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
				fmt.Printf("CTRL+C received. Stopping subscribers...\n")

				cancel()
			}
		}
	}()

	<-ctx.Done()

	fmt.Printf("program has been stopped\n")
}

func worker(ctx context.Context) {
	fmt.Printf("worker: start\n")

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker: stopped\n")

			return
		default:
			fmt.Printf("worker: running\n")

			time.Sleep(1 * time.Second)
		}
	}
}
