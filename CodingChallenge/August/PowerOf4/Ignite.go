package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPowerOfFour(16))
}

func isPowerOfFour(n int) bool {
	fmt.Println(n)

	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}
	if n%4 != 0 {
		return false
	}

	return isPowerOfFour(n / 4)
}
