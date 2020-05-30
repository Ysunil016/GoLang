package main

import (
	"fmt"
	"sort"
)

// Node ...
type Node struct {
	Val   int
	Coord []int
}

type sortByVal []Node

func (a sortByVal) Len() int           { return len(a) }
func (a sortByVal) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortByVal) Less(i, j int) bool { return a[i].Val < a[j].Val }

func main() {
	arr := [][]int{{1, 3}, {-2, 2}}
	fmt.Println(kClosest(arr, 1))
}

func kClosest(points [][]int, K int) [][]int {
	// Making Sorted Map Based on Distance from 0,0
	SortStore := make([]Node, len(points))
	for i := 0; i < len(points); i++ {
		Sum := points[i][0]*points[i][0] + points[i][1]*points[i][1]
		SortStore[i] = Node{Sum, points[i]}
	}
	sort.Sort(sortByVal(SortStore))
	Counter := 0
	// Now Getting Out K Nodes from Store.
	var result = make([][]int, K)

	for _, V := range SortStore {
		fmt.Println(V)
		result[Counter] = V.Coord
		Counter++
		if Counter >= K {
			break
		}
	}
	return result
}
