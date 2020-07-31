package main

import "fmt"

func main() {
	fmt.Println(climbStairs(5))
}

func climbStairs(n int) int {
	hash := make(map[int]int)
	return climb(n, 0, 0, hash)
}
func climb(n int, pos int, path int, hash map[int]int) int {
	_, V := hash[pos]
	if V {
		return hash[pos]
	}
	if pos > n {
		return 0
	}
	if pos == n {
		return 1
	}
	result := climb(n, pos+1, 1, hash) + climb(n, pos+2, 2, hash)
	hash[pos] = result
	return result
}
