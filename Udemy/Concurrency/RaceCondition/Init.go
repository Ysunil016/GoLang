package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	Counter := 0
	fmt.Println("CPUs ", runtime.NumCPU())

	gs := 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := Counter
			runtime.Gosched()
			v++
			Counter = v
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Routine ", runtime.NumGoroutine())
	fmt.Println(Counter)
}
