package main

import (
	"fmt"
	"sort"
)

func main() {
	Arr := [][]int{{7, 1}, {4, 4}, {7, 0}, {5, 0}, {6, 1}, {5, 2}}
	fmt.Println(reconstructQueue(Arr))
}

type sortByHeight [][]int
func (a sortByHeight) Len() int      { return len(a) }
func (a sortByHeight) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a sortByHeight) Less(i, j int) bool {
	if a[i][0] != a[j][0] {
		return a[i][0] > a[j][0]
	}
	return a[i][1] < a[j][1]
}

func reconstructQueue(people [][]int) [][]int {
	sort.Sort(sortByHeight(people))
	for I, V := range people {
		Value := people[I]
		for i := I; i > V[1]; i-- {
			people[i] = people[i-1]
		}
		people[V[1]] = Value
	}
	return people
}
