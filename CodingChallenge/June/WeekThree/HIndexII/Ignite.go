package main

import "fmt"

func main() {
	Arr := []int{0, 1, 3, 4, 5, 6}
	fmt.Println(hIndex(Arr))
}

func hIndex(x []int) int {
	if len(x) == 0 {
		return 0
	}
	l := 0
	r := len(x) - 1
	for l <= r {
		m := l + (r-l)/2
		if x[m] < len(x)-m {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return len(x) - l
}
