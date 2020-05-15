package main

import "fmt"

func main() {
	arr := []int{5, -3, 5}
	fmt.Println(findMaxInCircularArray(arr))
}

func findMaxInCircularArray(arr []int) int {
	// Getting Max Contigou Element
	single := kadane(arr)

	// *-1 and Find the Max
	Sum := 0
	for i := 0; i < len(arr); i++ {
		Sum += arr[i]
		arr[i] *= -1
	}
	circularSum := kadane(arr) + Sum
	if circularSum > single && circularSum != 0 {
		return circularSum
	}
	return single
}

func kadane(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	soFarMax := nums[0]
	endOfMax := 0

	for i := 0; i < len(nums); i++ {
		endOfMax += nums[i]
		if soFarMax < endOfMax {
			soFarMax = endOfMax
		}
		if endOfMax < 0 {
			endOfMax = 0
		}
	}
	return soFarMax
}
