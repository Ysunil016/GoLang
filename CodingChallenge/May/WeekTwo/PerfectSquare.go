package main

import "fmt"

func main() {
	fmt.Println(isPerfectSquare(16))
}

func isPerfectSquare(num int) bool {
	// Find the Point Where key is Greater then Assumed
	Counter := 1
	Product := 1
	for Product < num {
		Product = Counter * Counter
		Counter++
	}
	if Product == num {
		return true
	}
	return false
}
