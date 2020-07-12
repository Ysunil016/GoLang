package main

import "fmt"

func main() {
fmt.Println(reverseBits(266))	
}
func reverseBits(n uint32) uint32 {
	var res uint32 = 0
	t := 31
	for t > 0 {
		res |= (n&1)
		n = n >>1
		res = res <<1
		t=t-1
	}
	res |= (n&1)
	return res
}