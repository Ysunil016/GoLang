package main

import (
	"fmt"
)

func main() {
	var array = []int{1, 2, 1, 2, 3, 4, 5, 4}
	fmt.Println(singleNumber(array))
}

func singleNumber(nums []int) []int {
	var result = make([]int, 2)
	var difference int
	for _, x := range nums {
		difference ^= x
	}
	difference &= -difference
	for _, x := range nums {
		if x&difference == 0 {
			result[0] ^= x
		} else {
			result[1] ^= x
		}
	}
	return result
}
