package main

import "fmt"

func main() {
	fmt.Println(myPow(2, 5))
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	holding := myPow(x, n/2)
	if n%2 == 0 {
		return holding * holding
	}
	if n > 0 {
		return x * holding * holding
	}
	return (holding * holding) / x
}
