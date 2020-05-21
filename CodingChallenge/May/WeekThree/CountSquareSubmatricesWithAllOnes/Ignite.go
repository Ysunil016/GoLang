package main

import (
	"fmt"
	"math"
)

func main() {
	arr := [][]int{{0, 1, 1, 1}, {1, 1, 1, 1}, {0, 1, 1, 1}}
	fmt.Println(countSquares(arr))
}

func countSquares(matrix [][]int) int {
	Count := 0
	for i := 0; i < len(matrix); i++ {
		Count += matrix[i][0]
	}
	for i := 1; i < len(matrix[0]); i++ {
		Count += matrix[0][i]
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] != 0 {
				matrix[i][j] = int(math.Min(float64(matrix[i-1][j-1]), math.Min(float64(matrix[i-1][j]), float64(matrix[i][j-1])))) + 1
				Count += matrix[i][j]
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			fmt.Print(matrix[i][j])
		}
		fmt.Println()
	}
	return Count
}
