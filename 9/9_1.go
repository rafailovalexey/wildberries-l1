package main

import "fmt"

/*
	№ 9 (1-ое решение)

	Разработать конвейер чисел.
	Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
*/

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	numbers := make(chan int, len(array))
	results := make(chan int, len(array))

	go func() {
		defer close(numbers)
		defer close(results)

		for _, value := range array {
			numbers <- value
			results <- value * 2
		}
	}()

	for result := range results {
		fmt.Println(result)
	}
}
