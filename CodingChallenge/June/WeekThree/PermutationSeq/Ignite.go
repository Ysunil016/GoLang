package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(getPermutation(3, 3))
}

func getPermutation(n int, k int) string {
	list := make([]int, 0)
	fact := make([]int, n)
	fact[0] = 1
	for i := 1; i < n; i++ {
		fact[i] = i * fact[i-1]
	}
	for i := 1; i <= n; i++ {
		list = append(list, i)
	}
	k--
	var res string
	for i := n - 1; i >= 0; i-- {
		index := k / fact[i]
		res += strconv.Itoa(list[index])
		remove(list, index)
		k = k % fact[i]
	}
	return res
}
func remove(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}
