package main

import "fmt"

func main() {
	channel := make(chan<- int, 1) // Send Only

	channel <- 42

	fmt.Println(<-channel) // This will not work as Channel is Send Only Not Receive
}
