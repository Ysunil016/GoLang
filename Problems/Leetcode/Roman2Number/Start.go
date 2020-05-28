package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(change2Val("DCCCC"))
}

func change2Val(romanNumber string) int {
	ValStore := map[byte]int{byte(76): 50, byte(68): 500, byte(77): 1000, byte(67): 100, byte(88): 10, byte(73): 1, byte(86): 5}
	result := 0
	var i int
	Array := strings.Split(romanNumber, "")
	for i = 0; i < len(Array)-1; i++ {
		if ValStore[romanNumber[i]] < ValStore[romanNumber[i+1]] {
			result -= ValStore[romanNumber[i]]
		} else {
			result += ValStore[romanNumber[i]]
		}
	}
	return result + ValStore[romanNumber[i]]
}
