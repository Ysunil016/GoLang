package main

import "fmt"

func main() {
	arr := []int{1}
	fmt.Println(change(5000, arr))
}

func change(amount int, coins []int) int {
	dynamic := make([]int, amount+1)
	dynamic[0] = 1
	for _, c := range coins {
		for i := c; i <= amount; i++ {
			dynamic[i] += dynamic[i-c] // From Its Remainder
		}
	}
	return dynamic[amount]
}
