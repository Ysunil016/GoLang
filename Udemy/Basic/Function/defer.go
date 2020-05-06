package main

import "fmt"

func main() {
	foo()
	defer last() // Executed at Last Exection of Main Function
	bar()
}

func foo() {
	fmt.Println("Foo")
}
func bar() {
	fmt.Println("Bar")
	defer last() // Called at Last in bar Function
}

func last() {
	fmt.Println("Last")
}
