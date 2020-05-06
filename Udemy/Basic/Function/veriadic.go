package main

import "fmt"

func main() {
	Sum := foo(1, 2, 3, 4, 5, 6, 7, 8, 9)

	XS := []int{10, 20, 30, 40, 50}
	Sum2 := foo(XS...) // Unwrapping of XS...

	Sum3 := foo() // Nothing Passed also Works

	fmt.Println(Sum)
	fmt.Println(Sum2)
	fmt.Println(Sum3)
}

// Veriadic -- Unlimited Number of Arguments
func foo(x ...int) int { // Slice of Int
	fmt.Println(x)
	fmt.Printf("\n%T\n", x)
	var sum int = 0
	for _, v := range x {
		sum += v
	}
	return sum
}
