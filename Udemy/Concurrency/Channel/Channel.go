package main

import (
	"fmt"
	"time"
)

func getAction(X int) int {
	for i := 0; i < 10; i++ {
		time.Sleep(500 * time.Millisecond)
		X++
	}
	return X * 5
}

func main() {
	ch := make(chan int)
	ch2 := make(chan int)
	go func() {
		ch <- getAction(5)
	}()
	go func() {
		ch2 <- getAction(10)
	}()

	fmt.Println(<-ch)
	fmt.Println(<-ch2)
}
