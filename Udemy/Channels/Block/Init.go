package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int)

	channel <- 43
	time.Sleep(time.Second * 5)
	fmt.Println(<-channel)
}
