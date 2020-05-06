package main

import "fmt"

func main() {
	te()
}

func te() {
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(true && true)
	fmt.Println(false || false)
	fmt.Println(!false)
}
