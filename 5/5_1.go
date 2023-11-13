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
	№ 5 ( 1решение)

	Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
	По истечению N секунд программа должна завершаться.
*/

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	channel := make(chan string)

	go func() {
		fmt.Println("Subscriber: Start")

		for {
			select {
			case <-ctx.Done():
				fmt.Println("Subscriber: Stopped")

				return
			case message, isOpen := <-channel:
				if !isOpen {
					fmt.Println("Subscriber: Channel closed")

					return
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
