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
	№ 5 (1-ое решение)

	Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
	По истечению N секунд программа должна завершаться.
*/

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	channel := make(chan string)

	go func() {
		fmt.Printf("subscriber: start\n")

		for {
			select {
			case <-ctx.Done():
				fmt.Printf("subscriber: stopped\n")

				return
			case message, isOpen := <-channel:
				if !isOpen {
					fmt.Printf("subscriber: channel closed\n")

					return
				}

				fmt.Printf("subscriber: %s\n", message)
			}
		}
	}()

	go func() {
		defer close(channel)

		for counter := 0; ; counter++ {
			select {
			case <-ctx.Done():
				fmt.Printf("publisher: stopped\n")

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
