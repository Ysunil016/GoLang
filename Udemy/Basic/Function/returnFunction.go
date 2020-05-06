package main

import "fmt"

func main() {
	val, name := reFunction(454, "Sunil")()
	fmt.Printf("%v and %T\n", val, name)
}

func reFunction(val int, name string) func() (int, string) {
	funcT := func() (int, string) {
		return val, name
	}
	return funcT
}
