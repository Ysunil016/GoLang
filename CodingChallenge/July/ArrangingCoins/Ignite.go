package main

import "fmt"

func main() {
	fmt.Println(findCoin(5))
}

func findCoin(n int) int{
		Count := 0
		Res := 1
		for (n - Res) >= 0 {
			n -= Res;
			Res++
			Count++
		}
		return Count;
}