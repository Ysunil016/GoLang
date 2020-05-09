package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// Resource ... Atomic Sharable Data

func main() {

	var Resource int64
	var wg sync.WaitGroup
	const gs int = 100
	wg.Add(gs)

	fmt.Println("CPUs ", runtime.NumCPU())

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&Resource, 1)
			runtime.Gosched()
			fmt.Println("Res ", atomic.LoadInt64(&Resource))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(Resource)
}
