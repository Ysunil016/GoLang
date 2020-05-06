package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	res := logic(sum, arr...)
	fmt.Println(res)
}

func sum(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

func logic(f func(...int) int, nums ...int) int {
	var even []int
	for _, v := range nums {
		if v%2 == 0 {
			even = append(even, v)
		}
	}
	return f(even...)
}
