package main

import "fmt"

func main() {
	// One()
	// Two()
	// Three()
	// Four()
	// Five()
}

/////////////////////////////////////////////////////////////////////////////////////

// Five ... HandsOn Excersice 5
func Five() {
	type XType string
	type YType XType

	var X5 XType
	var Y5 YType

	fmt.Printf("%v\t%T\n", X5, X5)
	X5 = "Sunil Here"
	fmt.Printf("%v\t%T\n", X5, X5)
	var fResultX string = string(X5)
	fmt.Printf("%v\t%T\n", fResultX, fResultX)
	Y5 = "Sanjay Here"
	fmt.Printf("%v\t%T\n", Y5, Y5)
	var fResultY string = string(Y5)
	fmt.Printf("%v\t%T\n", fResultY, fResultY)

}

/////////////////////////////////////////////////////////////////////////////////////

// Four ... HandsOn Excersice 4
func Four() {

	type Sunil int // Own Type Created
	var x4 Sunil
	fmt.Printf("%v\t%T\n", x4, x4)
	x4 = Sunil(42)
	fmt.Printf("%v\t%T\n", x4, x4)
}

/////////////////////////////////////////////////////////////////////////////////////

var x3 int = 42
var y3 string = "James Bond"
var z3 = true

// Three ... HandsOn Excersice 3
func Three() {
	Spring := fmt.Sprintf("%+v\t%+v\t%+v", x3, y3, z3)
	fmt.Println(Spring)
}

/////////////////////////////////////////////////////////////////////////////////////

var x2 int
var y2 string
var z2 bool

// Two ... HandsOn Excersice 2
func Two() {
	fmt.Println(x2)
	fmt.Println(y2)
	fmt.Println(z2)
}

/////////////////////////////////////////////////////////////////////////////////////

// One ... HandsOn Excersice 1
func One() {
	x := 42
	y := "James Bond"
	z := true

	//Single Line Print
	fmt.Println(x, ",", y, ",", z)

	// MultiLine Print
	fmt.Printf("%+v\n", x)
	fmt.Printf("%+v\n", y)
	fmt.Printf("%+v\n", y)
}
