package main

import "fmt"

func main() {
	declareArray()
}

func declareArray() {
	var X2 [10]int // Zero Value
	fmt.Println(X2)
	X2[3] = 69
	fmt.Println(X2)
	fmt.Println(len(X2)) //Len Array

	// X1:= {1,2,3,4,5}
	// fmt.Println(X1)

}
