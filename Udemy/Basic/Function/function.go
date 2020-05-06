package main

import "fmt"

func main() {
	foo()
	foo2("Sunil")
	fmt.Println(add(1, 2))
	X, Y := mulTask("Sunil", "Yadav")
	fmt.Println(X, Y)
}

// func (r receiver) fname(arg type...) return { }

func foo() {
	fmt.Println("Sunil")
}

// Everything is Pass by Value
func foo2(s string) {
	fmt.Println(s)
}

func add(a int, b int) int {
	return a + b
}

func mulTask(fname string, lname string) (string, bool) {
	return fname, true
}
