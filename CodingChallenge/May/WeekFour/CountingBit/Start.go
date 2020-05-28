package main

import "fmt"

func main() {
	fmt.Println(makeArray(16))
}

func makeArray(nums int) []int {
	arr := make([]int, nums)
	arr[0] = 0

	// Handling Special Case
	if nums == 0 {
		return arr
	}
	arr[1] = 1
	if nums == 1 {
		return arr
	}
	// Pattern Recognized
	for i := 2; i < nums; i++ {
		if i%2 == 0 {
			arr[i] = arr[i/2]
		} else {
			arr[i] = arr[i/2] + 1
		}
	}
	return arr
}
