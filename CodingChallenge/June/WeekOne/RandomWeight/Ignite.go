package main

import (
	"fmt"
	"math/rand"
)

func main() {
	Arr := []int{1, 2, 3, 2, 3}
	Solution := Constructor(Arr)
	for i := 0; i < len(Arr); i++ {
		fmt.Print(Solution.PickIndex())
	}
}

// Solution ...
type Solution struct {
	Total        float32
	WeightPrefix []int
}

// Constructor ...
func Constructor(w []int) Solution {
	WeightPrefix := make([]int, len(w))
	TSum := 0
	for i := 0; i < len(w); i++ {
		TSum += w[i]
		WeightPrefix[i] += TSum
	}
	return Solution{float32(TSum), WeightPrefix}
}

// PickIndex ...
func (sol *Solution) PickIndex() int {
	MathRandom := rand.Float32() * sol.Total
	low := 0
	high := len(sol.WeightPrefix)
	for low < high {
		mid := (high + low) / 2
		if MathRandom > float32(sol.WeightPrefix[mid]) {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}
