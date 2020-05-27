package main

import "fmt"

func main() {
	arr := [][]int{{1, 2}, {1, 3}, {2, 4}}
	fmt.Println(possibleBipartition(4, arr))
}

func possibleBipartition(N int, dislikes [][]int) bool {
	// Init Visited Array First
	visited := make([]int, N)
	for i := 0; i < N; i++ {
		visited[i] = -1
	}

	// Make AdjList List
	adjList := make(map[int][]int)
	for i := 0; i < len(dislikes); i++ {
		adjList[dislikes[i][0]-1] = append(adjList[dislikes[i][0]-1], dislikes[i][1]-1)
		adjList[dislikes[i][1]-1] = append(adjList[dislikes[i][1]-1], dislikes[i][0]-1)
	}

	// DFS Search for the Node Having Same GROUP Symbol
	for i := 0; i < N; i++ {
		if visited[i] == -1 && !dfs(adjList, visited, i, 0) {
			return false
		}
	}

	return true
}

func dfs(adjList map[int][]int, visited []int, index int, grp int) bool {
	// If Neighbour was Visited then Check If If Is It Going for Same Group or Not
	if visited[index] == -1 {
		visited[index] = grp
	} else {
		return visited[index] == grp
	}

	// Visting Neighbour of Current Node... In Depth
	for _, V := range adjList[index] {
		if !dfs(adjList, visited, V, 1-grp) {
			return false
		}
	}
	return true
}
