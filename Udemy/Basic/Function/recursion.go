package main

import "fmt"

func main() {
	fmt.Println(fact(6))
}

func fact(num int) int {
	if num == 1 {
		return 1
	}
	return num * fact(num-1)
}
