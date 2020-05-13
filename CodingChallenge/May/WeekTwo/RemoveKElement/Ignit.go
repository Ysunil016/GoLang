package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "10200"
	fmt.Println(removeKdigits(input, 1))
}
func removeKdigits(num string, k int) string {
	var stack []string
	for i := 0; i < len(num); i++ {
		for k > 0 && len(stack) != 0 && stack[(len(stack)-1)] > string(num[i]) {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, string(num[i]))
	}
	for k > 0 {
		stack = stack[:len(stack)-1]
		k--
	}
	var Result []string

	fmt.Println(stack)
	for i := 0; i < len(stack); i++ {
		Result = append(Result, stack[i])
	}
	fmt.Println(Result)
	for len(Result) > 0 && Result[0] == "0" {
		Result = Result[1:]
	}
	if len(Result) == 0 {
		return "0"
	}
	return strings.Join(Result, "")
}
