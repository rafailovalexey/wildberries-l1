package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	channel := make(chan string)
	wg := &sync.WaitGroup{}

	wg.Add(1)

	go func() {
		defer wg.Done()

		fmt.Println("Subscriber: Start")

		for {
			select {
			case <-ctx.Done():
				fmt.Println("Subscriber: Stopped")

				return
			case message, isOpen := <-channel:
				if !isOpen {
					fmt.Println("Subscriber: Channel closed")
				}

				fmt.Printf("Subscriber: %s\n", message)
			}
		}
	}()

	go func() {
		defer close(channel)

		for counter := 0; ; counter++ {
			select {
			case <-ctx.Done():
				fmt.Println("Publisher: Stopped")

				return
			default:
				message := fmt.Sprintf("%d", counter)
				channel <- message
				counter++
			}
		}
	}()

	go func() {
		exit := make(chan os.Signal)
		signal.Notify(exit, syscall.SIGINT)
		<-exit

		fmt.Println("CTRL+C received. Stopping subscribers...")

		cancel()
		wg.Wait()

		fmt.Println("Program has been stopped")
	}()

	wg.Wait()
}
