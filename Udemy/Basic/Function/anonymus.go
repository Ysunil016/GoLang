package main

import "fmt"

func main() {

	// Auto Run Anonymus  Function Expression
	func() {
		fmt.Println("Pure Anonymus ")
	}()

	// Anonymus Function Expression
	X := func(v int) {
		fmt.Println("Anonymus Expression ", v)
	}
	X(44)

}
