package main

import "fmt"

var numStore map[string]int

func init() {
	numStore = map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
}

func main() {
	fmt.Println(convert2Roman(702))
}

func convert2Roman(num int) string {
	var Result string

	if num/1000 >= 1 {
		num -= 1000 * int(num/1000)
	} else if num/500 >= 1 {
		num -= 500 * int(num/500)
		fmt.Println(num)
	} else if num/100 >= 1 {
		num -= 100 * int(num/100)
	} else if num/50 >= 1 {
		fmt.Println(100)
		num -= 100 * int(num/50)
	} else if num/10 >= 1 {
		fmt.Println(10)
		num -= 10 * int(num/10)
	} else if num/5 >= 1 {
		fmt.Println(5)
		num -= 5 * int(num/5)
	} else if num/1 >= 1 {
		num--
	}

	fmt.Println(num)

	return Result
}
