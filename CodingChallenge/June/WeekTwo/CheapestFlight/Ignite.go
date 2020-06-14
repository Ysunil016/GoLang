package main

import (
	"fmt"
)

func main() {
	Arr := [][]int{{0, 1, 100}, {0, 2, 500}, {1, 2, 100}}
	fmt.Println(findCheapestPrice(3, Arr, 0, 2, 1))
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	MaxVal := 1 << 30
	DP := make([][]int, K+2)
	for i := 0; i < K+2; i++ {
		DP[i] = make([]int, n)
		for j := 0; j < n; j++ {
			DP[i][j] = MaxVal
		}
	}
	DP[0][src] = 0
	for i := 1; i <= K+1; i++ {
		DP[i][src] = 0
		for _, P := range flights {
			DP[i][P[1]] = min(DP[i][P[1]], DP[i-1][P[0]]+P[2])
		}
		fmt.Println(DP)
	}

	if DP[K+1][dst] >= MaxVal {
		return -1
	}
	return DP[K+1][dst]
}
func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
