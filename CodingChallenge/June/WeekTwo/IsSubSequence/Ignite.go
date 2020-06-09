package main

import "fmt"

func main() {
	fmt.Println(isSubsequence("abc", "asddsbsakjdbc"))
}

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(t) == 0 {
		return false
	}
	counter := 0
	for _, v := range []byte(t) {
		if s[counter] == v {
			counter++
		}
		if counter == len(s) {
			return true
		}
	}
	return false
}
