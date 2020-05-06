package main

import (
	"fmt"
	"sort"
)

func main() {
	X := []int{2, 2, 1, 1, 1, 2, 2}
	sort.Ints(X)
	fmt.Println(X[len(X)/2])
}
