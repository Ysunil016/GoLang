package main

import "fmt"

func main() {
	Arr := []int{0, 1, 2}
	fmt.Println(findMin(Arr))
}
func findMin2(nums []int) int {
	result := nums[0]
	for _, V := range nums {
		if result < V {
			result = V
		}
	}
	return result
}

func findMin(nums []int) int {
	low := 0
	high := len(nums) - 1
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] == nums[high] {
			high--
		} else if nums[mid] < nums[high] {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return nums[high]
}
