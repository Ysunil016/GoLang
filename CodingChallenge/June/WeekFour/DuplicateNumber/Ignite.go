package main

import "fmt"

func main() {
	fmt.Println(findDuplicate([]int{1, 2, 3, 2, 5}))
}
func findDuplicate(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return nums[i]
			}
		}
	}
	return -1
}
