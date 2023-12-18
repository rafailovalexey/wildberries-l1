package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
)

/*
	№ 4 (1-ое решение)

	Реализовать постоянную запись данных в канал (главный поток).
	Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
	Необходима возможность выбора количества воркеров при старте.
	Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/

/*
	Обоснование способа завершения работы всех воркеров.

	context предоставляет механизм отмены операций и передачи сигналов завершения.
	sync.WaitGroup используется для ожидания завершения всех воркеров.
	graceful shutdown используется для отлова сигнала CTRL+C

	Отлавливаем сигнал.
	После чего происходит запуск отмены контекста во всех воркерах.
	Затем дожидаемся всех оставшихся воркеров с помощью wait group.

	Все эти пункты позволяют безопасно дождаться завершения всех goroutines и сделать определённые операции перед завершением программы (закрыть коннект с базой данных допустим)
*/

func main() {
	workers := runtime.NumCPU()

	wg := &sync.WaitGroup{}
	channel := make(chan string)
	ctx, cancel := context.WithCancel(context.Background())

	for index := 1; index <= workers; index++ {
		wg.Add(1)

		log.Printf("subscriber-%d: started\n", index)

		go RunSubscriber(ctx, index, wg, channel)
	}

	go func() {
		defer close(channel)

		for counter := 0; ; counter++ {
			select {
			case <-ctx.Done():
				return
			default:
				message := log.Sprintf("publisher-1: %d\n", counter)

				channel <- message

				counter++
			}
		}
	}()

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT)
	<-exit

	log.Printf("ctrl+c received. Stopping subscribers...\n")

	cancel()
	wg.Wait()

	log.Printf("program has been stopped\n")
}

func RunSubscriber(ctx context.Context, index int, wg *sync.WaitGroup, channel <-chan string) {
	defer wg.Done()

	log.Printf("subscriber-%d: started\n", index)

	for {
		select {
		case <-ctx.Done():
			log.Printf("subscriber-%d: stopped\n", index)

			return
		case message, isOpen := <-channel:
			if !isOpen {
				log.Printf("subscriber-%d: channel closed\n", index)

				return
			}

			log.Printf("subscriber-%d: %s\n", index, message)
		}
	}
}
