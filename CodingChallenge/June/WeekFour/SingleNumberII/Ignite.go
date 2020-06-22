package main

import "fmt"

func main() {
	Arr := []int{2, 2, 2, 5}
	fmt.Println(singleNumber(Arr))
}

func singleNumber(nums []int) int {
	One := 0
	Two := 0
	var CommonBit int
	for i := 0; i < len(nums); i++ {
		Two = Two | (One & nums[i])
		One = One ^ nums[i]
		CommonBit = ^(One & Two) // Not Operator for Unirar Operation
		One &= CommonBit
		Two &= CommonBit
	}
	return One
}
