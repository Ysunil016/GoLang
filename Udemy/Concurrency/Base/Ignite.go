package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func init() {
	fmt.Println("Init")
}

var wg sync.WaitGroup

func main() {

	wg.Add(2) // Telling there is 2 Go Routines
	go foo()
	go bar()
	// fmt.Println("OS ", runtime.GOOS)
	// fmt.Println("ARCH ", runtime.GOARCH)
	// fmt.Println("CPU ", runtime.NumCPU())
	fmt.Println("Routines ", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("CPU ", runtime.NumCPU())
}

func foo() {
	for i := 0; i < 10000000000; i++ {
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10000000000; i++ {
	}
	wg.Done()
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
