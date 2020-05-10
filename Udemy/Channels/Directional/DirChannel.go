package main

import "fmt"

func main() {
	c:=make(chan int)
	cr:=make(<-chan int)
	cs:=make(chan <-int)

	fmt.Printf("%T\t",c)
	fmt.Printf("%T\t",cr)
	fmt.Printf("%T\t",cs)

}