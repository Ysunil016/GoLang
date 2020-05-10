package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	fmt.Printf("%v\n", ctx)
	fmt.Println(ctx.Err())
	fmt.Printf("%T\n", ctx)

	ctx, cancel := context.WithCancel(ctx)
	fmt.Printf("%v\n", ctx)
	fmt.Println(ctx.Err())
	fmt.Printf("%T\n", ctx)
	fmt.Printf("%v\n", cancel)
	fmt.Printf("%T\n", cancel)

	cancel()
	fmt.Println()
	fmt.Printf("%v\n", ctx)
	fmt.Println(ctx.Err())
	fmt.Printf("%T\n", ctx)
	fmt.Printf("%v\n", cancel)
	fmt.Printf("%T\n", cancel)

}
