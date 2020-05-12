package main

import "fmt"

func main() {
	arr := []int{1}
	fmt.Println(singleNonDuplicateS(arr))
}

// Optamized to O(lngn)
func singleNonDuplicateS(nums []int) int {
	return findElement(nums, 0, len(nums)-1)
}

func findElement(nums []int, start int, end int) int {
	fmt.Println(start, end)
	if start > end {
		return 0
	}
	mid := start + (end-start)/2

	if mid == 0 {
		return nums[mid]
	}
	if mid == len(nums)-1 {
		return nums[mid]
	}

	if nums[mid] != nums[mid-1] && nums[mid] != nums[mid+1] {
		return nums[mid]
	}

	if nums[mid] == nums[mid-1] {
		if mid%2 == 0 {
			return findElement(nums, start, mid)
		}
		return findElement(nums, mid+1, end)
	}
	if mid%2 == 0 {
		return findElement(nums, mid+1, end)
	}
	return findElement(nums, start, mid)

}

// This is O(n) time
func singleNonDuplicate(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}
