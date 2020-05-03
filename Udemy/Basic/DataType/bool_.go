package main

import "fmt"

var booleanValue bool

func main() {
	fmt.Println(booleanValue)
	booleanValue = true
	fmt.Println(booleanValue)
	fmt.Println(isGreater(24, 20))
	fmt.Println(isGreater(20, 24))
	fmt.Println(isGreater(20, 20))

	// There is No === Equal Operator in Golang
}

func isGreater(a int, b int) bool {
	return a > b
}
