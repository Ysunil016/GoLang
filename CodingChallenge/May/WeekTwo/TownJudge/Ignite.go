package main

import "fmt"

func main() {
	arr := [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}}
	fmt.Println(findJudge(4, arr))
}

func findJudge(N int, trust [][]int) int {
	PeopleStore := make([]int, N)
	for i := 0; i < len(trust); i++ {
		A := trust[i][0]
		B := trust[i][1]

		PeopleStore[A-1]--
		PeopleStore[B-1]++
	}
	fmt.Println(PeopleStore)

	hasJudge := false
	judgeIndex := -1
	for i := 0; i < N; i++ {
		if PeopleStore[i] == N-1 {
			if hasJudge == true {
				return -1
			} else {
				hasJudge = true
				judgeIndex = i + 1
			}
		}
	}
	return judgeIndex
}
