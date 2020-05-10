package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		N := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Cancled Context")
				return // Prevent Routine from Leaking
			default:
				N++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("Working ", N)
			}
		}
	}()

	///
	fmt.Println(ctx)
	fmt.Println(ctx.Err())
	///
	time.Sleep(time.Second * 2)

	fmt.Println(ctx)
	fmt.Println(ctx.Err())

	time.Sleep(time.Second * 2)
	cancel()

	time.Sleep(time.Second * 2)

	fmt.Println(ctx)
	fmt.Println(ctx.Err())

	// defer cancel() // Executes when Main is About to Terminate
}
