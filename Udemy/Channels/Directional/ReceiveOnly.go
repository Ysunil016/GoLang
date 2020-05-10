package main

import "fmt"

func main() {
	channel := make(<-chan int, 1) // Receive Only

	channel <- 42 /// This will not work as Channel is Receive Only Not Send

	fmt.Println(<-channel)
}
