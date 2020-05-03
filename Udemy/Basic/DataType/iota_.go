package main

import "fmt"

const (
	a = iota
	b = iota
	c
)
const (
	e = iota
	f
	g
)

func main() {
	playIota()
}

func playIota() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
}
