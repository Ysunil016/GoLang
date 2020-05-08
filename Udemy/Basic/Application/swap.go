package main

import "fmt"

func main() {
	X := 2
	Y := 3
	fmt.Println(X, Y)
	X, Y = Y, X
	fmt.Println(X, Y)
}
