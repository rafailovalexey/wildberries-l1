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
	№ 6 (3 решение)

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
				fmt.Println("CTRL+C received. Stopping subscribers...")

				cancel()
			}
		}
	}()

	<-ctx.Done()

	fmt.Println("Program has been stopped")
}

func worker(ctx context.Context) {
	fmt.Println("Worker: Start")

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker: Stopped")
			return
		default:
			fmt.Println("Worker: Running")
			time.Sleep(1 * time.Second)
		}
	}
}
