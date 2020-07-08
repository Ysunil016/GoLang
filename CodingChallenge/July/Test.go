package main

import "fmt"

func main() {
	var a int = 0x50
	var b float64 = 7.4
	var c float64

	c = float64(a) + float64(b)

	fmt.Println(c)
}
