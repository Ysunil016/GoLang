package main

import "fmt"

func main() {
	Arr := []int{2, 1, 2, 1, 2, 1, 0, 1, 0, 0, 0, 2, 1, 0, 2, 1}
	sortColors(Arr)
	for _, V := range Arr {
		fmt.Print(V)
	}
}

func sortColors(nums []int) {
	Start := 0
	End := len(nums) - 1
	Index := 0
	for Index <= End && Start < End {
		if nums[Index] == 0 {
			nums[Index] = nums[Start]
			nums[Start] = 0
			Start++
			Index++
		} else if nums[Index] == 2 {
			nums[Index] = nums[End]
			nums[End] = 2
			End--
		} else {
			Index++
		}
	}

}
