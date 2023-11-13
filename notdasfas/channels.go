package main

func main() {

}

//func main() {
//	channel := make(chan int)
//
//	go func() {
//		defer close(channel)
//
//		channel <- 1
//		channel <- 2
//		channel <- 3
//		channel <- 4
//		channel <- 5
//	}()
//
//	fmt.Println(<-channel)
//	fmt.Println(<-channel)
//	fmt.Println(<-channel)
//	fmt.Println(<-channel)
//	fmt.Println(<-channel)
//}

//func main() {
//	channel := make(chan int)
//
//	go func() {
//		fmt.Println(<-channel)
//		fmt.Println(<-channel)
//		fmt.Println(<-channel)
//		fmt.Println(<-channel)
//		fmt.Println(<-channel)
//	}()
//
//	channel <- 1
//	channel <- 2
//	channel <- 3
//	channel <- 4
//	channel <- 5
//
//	close(channel)
//}
