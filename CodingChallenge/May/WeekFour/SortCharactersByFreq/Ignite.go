package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(frequencySort("tree"))
}

func frequencySort(s string) string {
	store := make(map[string]int, 0)
	arr := strings.Split(s, "") // Spliting String
	// Storing in Map
	for _, v := range arr {
		store[v]++
	}
	// Sorting In Descending Order
	sort.SliceStable(arr, func(i, j int) bool {
		if store[arr[i]] == store[arr[j]] {
			return arr[i] < arr[j]
		}
		return store[arr[i]] > store[arr[j]]
	})
	// Joining the Splitted Parts
	return strings.Join(arr, "")
}
