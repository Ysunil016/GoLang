package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(minDistance("sea", "eat"))
}

func minDistance(X string, Y string) int {
	Word1 := strings.Split(X, "")
	Word2 := strings.Split(Y, "")
	Result := make([][]int, len(Word1)+1)

	for i := 0; i <= len(Word1); i++ {
		Result[i] = make([]int, len(Word2)+1)
	}

	for i := 0; i <= len(Word2); i++ {
		Result[0][i] = i
	}
	for i := 0; i <= len(Word1); i++ {
		Result[i][0] = i
	}

	for i := 1; i <= len(Word1); i++ {
		for j := 1; j <= len(Word2); j++ {
			if Word2[j-1] == Word1[i-1] {
				Result[i][j] = Result[i-1][j-1]
			} else {
				Result[i][j] = MinVal(MinVal(Result[i-1][j], Result[i][j-1]), Result[i-1][j-1]) + 1
			}
		}
	}
	return Result[len(Word1)][len(Word2)]
}

// MinVal ...
func MinVal(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
