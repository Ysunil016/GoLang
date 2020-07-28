package main

import (
	"fmt"
	"sort"
)

func main() {
	charBuffer := []byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}
	fmt.Println(leastInterval(charBuffer, 2))
}

type SortBy []int

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i] < a[j] }

func leastInterval(tasks []byte, n int) int {
	freqCount := make([]int, 26)
	for _, V := range tasks {
		freqCount[V-65]++
	}
	sort.Sort(SortBy(freqCount))

	maxElement := freqCount[25] - 1

	idealSlots := maxElement * n

	for i := 24; i >= 0; i-- {
		idealSlots -= min(freqCount[i], maxElement)
	}
	if idealSlots > 0 {
		return (idealSlots + len(tasks))
	}
	return len(tasks)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
