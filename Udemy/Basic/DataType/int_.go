package main

import (
	"fmt"
	"runtime"
)

func main() {
	usingInt()
}

func usingInt() {
	x := 42
	y := 50.2

	fmt.Printf("%v and %T\n", x, x)
	fmt.Printf("%v and %T\n", y, y)

	var uniX uint8 = 255
	fmt.Println(uniX)

	var iX int8 = -128
	fmt.Println(iX)

	// Getting Runtime Data
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

}
