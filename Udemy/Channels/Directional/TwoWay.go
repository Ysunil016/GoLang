package main

import "fmt"

func main() {
	channel := make(chan int)

	go foo(channel) // This will be Block as Seperate Routine Untill Bar Channel Receives It
	bar(channel)

	defer last()
}

func last() {
	fmt.Println("Program Exited")
}

/// Chan Sending Function
func foo(send chan<- int) {
	send <- 69
}

/// Chan Receiver Function
func bar(rec <-chan int) {
	fmt.Println(<-rec)
}
