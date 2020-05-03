package main

import "fmt"

const value int = 20000
const value2 = 20000.200
const value3 = "Sunil"
const value4 = "Rash"

func main() {
	fmt.Printf("%v of Type %T\n", value, value)
	fmt.Printf("%v of Type %T\n", value2, value2)
	fmt.Printf("%v of Type %T\n", value3, value3)
	fmt.Printf("%v of Type %T\n", value4, value4)
}
