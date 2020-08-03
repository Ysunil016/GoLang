package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isPalindrome("race is car"))
}

func isPalindrome(s string) bool {
	start := 0
	end := len(s) - 1
	s = strings.ToLower(s)
	for start < end {
		if (byte(s[start]) >= 48 && byte(s[start]) <= 57) || (byte(s[start]) >= 97 && byte(s[start]) <= 122) {
			if (byte(s[end]) >= 48 && byte(s[end]) <= 57) || (byte(s[end]) >= 97 && byte(s[end]) <= 122) {
				if s[start] != s[end] {
					return false
				} else {
					start++
					end--
				}
			} else {
				end--
			}
		} else {
			start++
		}
	}
	return true
}
