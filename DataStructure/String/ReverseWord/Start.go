package main

import (
	"fmt"
	"strings"
)

func main() {
	Word := "Sunil.Yadav.Is.Here"
	res := reverseWord(Word)
	fmt.Println(res)
}
func reverseWord(str string) string {
	arr := strings.Split(str, ".")
	var finalString string
	for i := len(arr) - 1; i >= 0; i-- {
		if i != 0 {
			finalString += arr[i] + "."
		} else {
			finalString += arr[i]
		}
	}
	return finalString
}
