package main

import (
	"fmt"
	"sync"
)

// var wg sync.WaitGroup

func main() {
	// Declaring Channels
	even := make(chan int)
	odd := make(chan int)
	fanIn := make(chan int) // Its a Design Pattern

	// wg.Add(1)

	// fmt.Printf("%T", wg)
	// Sending Inputs to Channel
	go send(even, odd)

	// Receive
	go receive(even, odd, fanIn)

	// FanIn Patter
	var wg sync.WaitGroup
	wg.Add(1)
	go func() { // This Will Restrict the Main Process from Stagnating...
		for v := range fanIn {
			fmt.Println("FanIn ", v)
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Exiting Program")

}

func receive(e, o <-chan int, fanIn chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for x := range e {
			fanIn <- x
		}
		wg.Done()
	}()
	go func() {
		for x := range o {
			fanIn <- x
		}
		wg.Done()
	}()
	wg.Wait()
	close(fanIn)
}

func send(e, o chan<- int) {
	for i := 0; i < 1000; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(o)
	close(e)
}
