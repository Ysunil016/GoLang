package main

import "fmt"

func main() {
	fmt.Println("Sunil Yadav")

	// for i := 1; i < 100; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Println(i)
	// 	}
	// }
	// base()
	declareShort()
}

func base() {
	n, _ := fmt.Println(1, 2, 3, 4, "Sunil", "Yadav", true)
	fmt.Println("Total Byte Size ", n)
}

func declareShort() {
	x := 42 // Short Declaration
	fmt.Println(x)
	x = 50
	fmt.Println(x)
	y := 100 + 24 // Expession
	fmt.Println(y)
	var z int // Default is 0, Lazy Declaration
	fmt.Println(z)
	z = 100
	fmt.Println(z)
}
