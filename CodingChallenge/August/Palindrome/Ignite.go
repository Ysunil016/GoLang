package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome(("abaaccaac")))
}

func longestPalindrome(s string) int {
	alphaArray := make([]int, 80)
	for _, c := range s {
		alphaArray[c-'A']++
	}
	result := 0
	for i := 0; i < 80; i++ {
		result += alphaArray[i] / 2
	}
	result *= 2
	if result < len(s) {
		result++
	}
	return result
}
