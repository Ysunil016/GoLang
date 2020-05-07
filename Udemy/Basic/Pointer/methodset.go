package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Print(s.area(), "\n")
}

func main() {

	C := circle{10}
	info(C)
	info(&C)

}
