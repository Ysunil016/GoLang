package main

import "fmt"

func main() {
	var X uint8 = 10 // 00001010
	fmt.Printf("%d with Binary %b\n", X, X)
	Y := shiftLeft(X, 2)
	fmt.Printf("%d with Binary %b\n", Y, Y) // 00101000

	fmt.Println(10 ^ 10 ^ 1) // XoR

	metric()
}

func shiftLeft(key uint8, counter int) uint8 {
	return key << counter
}

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func metric() {
	fmt.Println(kb)
	fmt.Println(mb)
	fmt.Println(gb)
}
