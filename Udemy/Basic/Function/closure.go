package main

import "fmt"

var i int

func main() {
	fmt.Println(i)
	i++
	fmt.Println(i)
	foo()
	fmt.Println(i)

	// Closure
	{
		y := 100
		fmt.Println(y)
	}
	// fmt.Println(y) // Y will Now be Accessable here, its scope is define in code block only

}

func foo() {
	fmt.Println(i)
}
