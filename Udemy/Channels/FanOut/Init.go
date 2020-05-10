package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go send(ch1)
	go tasking(ch1, ch2)

	counter := 0
	for <-ch2 != 0 { // Looping Over Ch2
		counter++
	}
	fmt.Println(counter)
}

func send(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

func tasking(ch <-chan int, ch2 chan<- int) {
	var wg sync.WaitGroup
	for e := range ch {
		wg.Add(1)
		go func(e int) {
			ch2 <- toHardWork(e)
			wg.Done()
		}(e)
	}
	wg.Wait()
	close(ch2)
}

func toHardWork(e int) int {
	time.Sleep(time.Millisecond * 1000)
	return 1000 + rand.Intn(1000)
}
