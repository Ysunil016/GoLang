package main

import (
	"fmt"
)

func main() {
	// Declaring Channels
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	// Sending Inputs to Channel
	go send(even, odd, quit)

	// Receive
	receive(even, odd, quit)
	fmt.Println("Exiting Program")

}

func receive(e, o <-chan int, q <-chan bool) {
	for {
		select {
		case v := <-e:
			fmt.Println(v, "is Even")
		case v := <-o:
			fmt.Println(v, "is Odd")
		case v, ok := <-q:
			fmt.Println(v, ok)
			return

		}
	}
}

func send(e, o chan<- int, q chan<- bool) {
	for i := 0; i < 1000; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(q)
}
