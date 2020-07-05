package main

import "fmt"

func main() {
	fmt.Println(hammingDistance(4, 1))
}
func hammingDistance(X int, Y int) int {
	t := 33
	r := 0
	for t != 0 {
		t--
		_X := X & 1
		_Y := Y & 1
		if _X != _Y {
			r++
		}
		X = X >> 1
		Y = Y >> 1
	}
	return r
}
