package main

//func main() {
//	array := make([]int, 0, 10)
//	array = append(array, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
//
//	numbers := make(chan int, len(array))
//	results := make(chan int, len(array))
//
//	go func() {
//		defer close(numbers)
//		defer close(results)
//
//		for _, value := range array {
//			numbers <- value
//			results <- value * 2
//		}
//	}()
//
//	for result := range results {
//		fmt.Println(result)
//	}
//}

//func main() {
//	array := make([]int, 0, 10)
//	array = append(array, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
//
//	numbers := make(chan int)
//	results := make(chan int)
//
//	go func() {
//		for result := range results {
//			fmt.Println(result)
//		}
//	}()
//
//	go func() {
//		for _ = range numbers {
//
//		}
//	}()
//
//	for _, value := range array {
//		numbers <- value
//		results <- value * 2
//	}
//
//	close(numbers)
//	close(results)
//}
