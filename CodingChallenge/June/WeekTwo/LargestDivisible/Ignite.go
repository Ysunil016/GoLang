package main

import (
	"fmt"
	"sort"
)

func main() {
	Arr := []int{1, 2, 4, 8, 16}
	for _, x := range largestDivisibleSubset(Arr) {
		fmt.Print(x)
	}
}

// SortNums ...
type SortNums []int

func (a SortNums) Len() int           { return len(a) }
func (a SortNums) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortNums) Less(i, j int) bool { return a[i] < a[j] }

func largestDivisibleSubset(nums []int) []int {
	sort.Sort(SortNums(nums))
	Result := make([][]int, len(nums))
	for X := range Result {
		Result[X] = append(Result[X], nums[X])
	}
	fmt.Println(Result)
	ResultX := []int{}
	for i, Value := range nums {
		for j := 0; j < i; j++ {
			if Value%nums[j] == 0 && len(Result[i]) < len(Result[j])+1 {
				fmt.Println(Result[j], Value)
				Result[i] = append(Result[j], Value)
			}
		}
	}
	fmt.Println(Result)

	for _, V := range Result {
		if len(V) > len(ResultX) {
			ResultX = V
		}
	}
	return ResultX
}
