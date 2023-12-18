package main

import "fmt"

/*
	№ 9 (2-ое решение)

	Разработать конвейер чисел.
	Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
*/

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	numbers := make(chan int)
	results := make(chan int)

	go func() {
		for result := range results {
			fmt.Printf("%d\n", result)
		}
	}()

	go func() {
		for _ = range numbers {
		}
	}()

	for _, value := range array {
		numbers <- value
		results <- value * 2
	}

	defer close(numbers)
	defer close(results)
}
