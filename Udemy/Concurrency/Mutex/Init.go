package main

import (
	"fmt"
	"runtime"
	"sync"
)

var mutex sync.Mutex

func main() {
	Counter := 0
	fmt.Println("CPUs ", runtime.NumCPU())
	gs := 10000
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mutex.Lock() // Locking Counter Variable for Usgage to Other Routine
			v := Counter
			runtime.Gosched()
			v++
			Counter = v
			mutex.Unlock() // UnLocking Counter Variable for Usgage to Other Routine
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Routine ", runtime.NumGoroutine())
	fmt.Println(Counter)
}
