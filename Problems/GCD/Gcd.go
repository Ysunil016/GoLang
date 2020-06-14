package main

import "fmt"

func main() {
	fmt.Println(recursiveGCD(30, 27))
	fmt.Println(iterativeGCD(30, 15))
}

func recursiveGCD(a int, b int) int {
	if b == 0 {
		return a
	}
	return recursiveGCD(b, a%b)
}

func iterativeGCD(a int, b int) int {
	for b != 0 {
		Temp := a % b
		a = b
		b = Temp
	}
	return a
}
