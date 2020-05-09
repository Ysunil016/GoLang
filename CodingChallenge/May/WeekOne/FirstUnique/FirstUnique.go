package main

import (
	"fmt"
)

func main() {
	fmt.Println(getUnique("abdefghabcd"))
}

func getUnique(s string) int {
	length := len(s)
	var arr [26]int
	for i := 0; i < length; i++ {
		arr[s[i]-'a']++
	}
	for i := 0; i < length; i++ {
		if arr[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}
