package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findComplement(5000))
}

func findComplement(num int) int {
	Bit := 1
	var bitsNeededToShift int = int(math.Log2(float64(num))) + 1
	for i := 0; i < bitsNeededToShift; i++ {
		num = num ^ Bit
		Bit = Bit << 1
	}
	return num
}
