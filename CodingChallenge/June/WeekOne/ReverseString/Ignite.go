package main

import "fmt"

func main() {
	Arr := []byte{'A', 'B', 'C', 'D'}
	fmt.Println(Arr)
	reverseString(Arr)
	fmt.Println(Arr)
}

func reverseString(s []byte) {
	for x := 0; x < len(s)/2; x++ {
		T := s[x]
		s[x] = s[len(s)-x-1]
		s[len(s)-x-1] = T
	}
}
