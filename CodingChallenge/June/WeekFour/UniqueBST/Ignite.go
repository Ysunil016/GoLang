package main

import "fmt"

func main() {
	fmt.Println(num(5))	
}

func num(x int)int{
	dp := make([]int,x+1)
	dp[0] = 1;
	dp[1] = 1;
	for i:=2;i<=x;i++{
		for j:=1;j<=i;j++{
			dp[i] = dp[i]  + (dp[i-j]*dp[j-1])
		}
	}
	return dp[x]
}