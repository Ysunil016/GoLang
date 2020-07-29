package main

func main() {

}

func maxProfit(prices []int) int {
	hashMap := make(map[string]int)
	return profit(prices, 0, 0, hashMap)
}

func profit(prices []int, isBuy int, rIndex int, hm map[string]int) int {
	if rIndex >= len(prices) {
		return 0
	}
	Key := string(rIndex) + string(isBuy)
	V, isAvailable := hm[Key]
	if isAvailable {
		return V
	}
	Result := 0
	if isBuy == 0 {
		isBuy := profit(prices, 1, rIndex+1, hm) - prices[rIndex]
		noBuy := profit(prices, 0, rIndex+1, hm)
		Result = max(isBuy, noBuy)
	} else {
		isSell := profit(prices, 0, rIndex+2, hm) + prices[rIndex]
		noSell := profit(prices, 1, rIndex+1, hm)
		Result = max(isSell, noSell)
	}

	hm[Key] = Result
	return Result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
