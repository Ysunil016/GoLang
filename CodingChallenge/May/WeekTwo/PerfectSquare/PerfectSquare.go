package main

import "fmt"

func main() {
	fmt.Println(isPerfectSquare2(808201))
}

func isPerfectSquare2(num int) bool {
	low := 1
	high := int(num / 2)
	mid := low + (high-low)/2

	mProd := mid * mid

	for mProd != num && low != high {
		if mProd > num {
			high = mid
			mid = low + (high-low)/2
			mProd = mid * mid
		} else {
			low = mid + 1
			mid = low + (high-low)/2
			mProd = mid * mid
		}
	}
	if mProd == num {
		return true
	}
	return false
}

func isPerfectSquare(num int) bool {
	// Find the Point Where key is Greater then Assumed
	Counter := 1
	Product := 1
	for Product < num {
		Product = Counter * Counter
		Counter++
	}
	if Product == num {
		return true
	}
	return false
}
