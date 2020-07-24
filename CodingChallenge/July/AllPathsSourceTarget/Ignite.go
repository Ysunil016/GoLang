package main

import "fmt"

func main() {
	arr := [][]int{{1, 2}, {3}, {3}, {}}
	allPathsSourceTarget(arr)
}

// Result ..
var Result [][]int
var indexCounter = 0

func allPathsSourceTarget(graph [][]int) [][]int {
	Result = make([][]int, len(graph))
	dfsSearch(0, len(graph)-1, 1, make([]int, len(graph)), graph)
	
	return Result
}

func dfsSearch(currentIndex, finalIndex, inCounter int, currentList []int, graph [][]int) {
	if currentIndex == finalIndex {
		fmt.Println(currentList)
		Result[indexCounter] = currentList
		return
	}

	for _, V := range graph[currentIndex] {
		currentList[inCounter] = V
		dfsSearch(V, finalIndex, inCounter+1, currentList, graph)
		currentList[inCounter] = -1
	}
}
