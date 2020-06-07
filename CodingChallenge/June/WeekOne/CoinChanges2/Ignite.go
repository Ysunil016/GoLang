package main

import "fmt"

func main() {
	arr := []int{1, 2, 5}
	fmt.Println(change(5, arr))
}

func change(amount int, coins []int) int {
	dynamic := make([]int, amount+1)
	dynamic[0] = 1
	for _, c := range coins {
		fmt.Println(c)
		for i := c; i <= amount; i++ {
			fmt.Println(dynamic)
			dynamic[i] += dynamic[i-c] // From Its Remainder
		}
	}
	return dynamic[amount]
}
