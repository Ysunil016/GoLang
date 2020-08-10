package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(titleToNumber("XYA"))
}

func titleToNumber(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		factor:=int(math.Pow(26,float64(i)))
		result += factor * int((byte(s[len(s)-1-i]) - 'A' +1))
	}
	return result
}
