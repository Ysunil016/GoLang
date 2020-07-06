package main

import (
	"fmt"
)

func main() {
	Arr := []int{1, 2, 4}
	fmt.Println(plusOne(Arr))
}

func plusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+carry == 10 {
			digits[i] = 0
		} else {
			digits[i] += carry
			carry = 0
			return digits
		}
	}
	nArr := make([]int, len(digits)+1)
	for i := len(digits) - 1; i >= 0; i-- {
		nArr[i+1] = digits[i]
	}
	nArr[0] = carry
	return nArr
}
