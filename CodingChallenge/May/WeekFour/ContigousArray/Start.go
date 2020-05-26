package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{0, 1}
	fmt.Println(findMaxLength(arr))
}

func findMaxLength(nums []int) int {
	result := 0
	counter := 0
	var store = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			counter--
		} else {
			counter++
		}
		if counter == 0 {
			result = i + 1
		}
		if _, ok := store[counter]; ok {
			result = int(math.Max(float64(result), float64(i-store[counter])))
		} else {
			store[counter] = i
		}
	}
	return result
}
