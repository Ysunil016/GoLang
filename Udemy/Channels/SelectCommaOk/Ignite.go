package main

import "fmt"

func main() {

	ch := make(chan int)
	go func() {
		ch <- 69
		close(ch) // If Not Closed <-ch will be in Deadlock, that waits for data on Channel
	}()

	val, ok := <-ch
	fmt.Println(val, ok) // Returns True with VAlue
	val, ok = <-ch       // Channel is Exhausted
	fmt.Println(val, ok) // Returns False, Since Channel is Closed
}
