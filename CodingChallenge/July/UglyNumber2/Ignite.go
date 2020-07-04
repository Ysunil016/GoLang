package main

import (
	"fmt"
)

func main() {
	fmt.Println(nthUglyNumber(30))
}

func nthUglyNumber(n int) int {
	if n == 1 {
		return 1
	}
	Store := make([]int, 0)
	Store = append(Store, 1)
	var i, j, k int

	for len(Store) < n {
		Store = append(Store, min(Store[i]*2, min(Store[j]*3, Store[k]*5)))
		if Store[i]*2 == Store[len(Store)-1] {
			i++
		}
		if Store[j]*3 == Store[len(Store)-1] {
			j++
		}
		if Store[k]*5 == Store[len(Store)-1] {
			k++
		}
	}
	return Store[len(Store)-1]
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
