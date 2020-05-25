package main

import (
	"log"
	"math"
)

func main() {
	A := []int{2, 5, 1, 2, 5}
	B := []int{10, 5, 2, 1, 5, 2}
	log.Println(maxUncrossedLines(A, B))

}

func maxUncrossedLines(A []int, B []int) int {
	dpArray := make([][]int, len(A))
	for i := range dpArray {
		dpArray[i] = make([]int, len(B))
	}
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			dpArray[i][j] = 0
		}
	}
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			if A[i] == B[j] {
				if i > 0 && j > 0 {
					dpArray[i][j] = dpArray[i-1][j-1] + 1
				} else {
					dpArray[i][j] = 1
				}
			} else {
				if i > 0 && j > 0 {
					dpArray[i][j] = int(math.Max(float64(dpArray[i-1][j]), float64(dpArray[i][j-1])))
				} else {
					if i == 0 && j != 0 {
						dpArray[i][j] = dpArray[i][j-1]
					}
					if i != 0 && j == 0 {
						dpArray[i][j] = dpArray[i-1][j]
					}
				}
			}
		}
	}
	return dpArray[len(A)-1][len(B)-1]
}
