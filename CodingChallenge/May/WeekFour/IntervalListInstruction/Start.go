package main

import (
	"fmt"
	"math"
)

func main() {
	arrA := [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}
	arrB := [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}
	result := intervalIntersection(arrA, arrB)
	fmt.Println(result)
}

func intervalIntersection(A [][]int, B [][]int) [][]int {
	result := [][]int{}
	i, j := 0, 0
	for i < len(A) && j < len(B) {
		start := int(math.Max(float64(A[i][0]), float64(B[j][0])))
		end := int(math.Min(float64(A[i][1]), float64(B[j][1])))
		if start <= end {
			result = append(result, []int{start, end})
		}
		if A[i][1] < B[j][1] {
			i++
		} else {
			j++
		}
	}
	return result
}
