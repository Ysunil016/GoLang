package main

import "fmt"

func main() {
	Arr := []int{1, 4, 6, 8}
	fmt.Print(searchInsert(Arr, 7))
}

func searchInsert(nums []int, target int) int {
	if target > nums[len(nums)-1] {
		return len(nums)
	}
	if target < nums[0] {
		return 0
	}
	low := 0
	high := len(nums)
	for low < high {
		mid := low + (high-low)/2
		if target > nums[mid] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}
