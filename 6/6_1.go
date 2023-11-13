package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go worker(ctx)

	exit := make(chan os.Signal)
	signal.Notify(exit, syscall.SIGINT)
	<-exit

	fmt.Println("CTRL+C received. Stopping subscribers...")

	cancel()

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
