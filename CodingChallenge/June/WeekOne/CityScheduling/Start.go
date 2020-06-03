package main

import (
	"fmt"
	"sort"
)

func main() {
	Array := [][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}}
	fmt.Println(twoCitySchedCost(Array))
}

// SortNumber ...
type SortNumber [][]int
func (a SortNumber) Len() int           { return len(a) }
func (a SortNumber) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortNumber) Less(i, j int) bool { return (a[i][0] - a[i][1]) < (a[j][0] - a[j][1]) }
func twoCitySchedCost(costs [][]int) int {
	sort.Sort(SortNumber(costs))
	fmt.Println(costs)
	Result := 0
	for I := 0; I < len(costs); I++ {
		if I < len(costs)/2 {
			Result += costs[I][0]
		} else {
			Result += costs[I][1]
		}
	}
	return Result
}
