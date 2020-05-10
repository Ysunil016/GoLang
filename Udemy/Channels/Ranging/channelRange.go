package main

import "fmt"

func main() {
	channel := make(chan int)

	// Pusing Data to Channel
	go func() {
		for i := 0; i < 100; i++ {
			channel <- i
		}
		close(channel) // Closing Channel Will Only Allow the Ranging to Stop
	}()

	// Reading Channel Data using Range Clause
	for v := range channel {
		fmt.Println(v)
	}

	fmt.Println("Exiting Main")

}
