package main

import "fmt"

func main() {
	fmt.Println(isPowerOf2(128))
}

func isPowerOf2(num int) bool {
	if num < 0 {
		return false
	}
	count := 0
	for num != 0 {
		if num%2 != 0 {
			count++
		}
		num = num >> 1
	}
	if count == 1 {
		return true
	}
	return false
}
