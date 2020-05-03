package main

import "fmt"

func main() {
	fmtView()
	customType()
}

// fmtView ... This will give the way to present values inside fmt
func fmtView() {
	X := 911

	fmt.Printf("%T\n", X)  // Data Type
	fmt.Printf("%b\n", X)  // Binary Base 2
	fmt.Printf("%d\n", X)  // Decimal Base 10
	fmt.Printf("%o\n", X)  // Octal Base 8
	fmt.Printf("%x\n", X)  // Hex SMALL Base 16
	fmt.Printf("%#X\n", X) // Hex CAP Base 16
	fmt.Printf("%c\n", X)  // Character at Unicode Point
	fmt.Printf("%U\n", X)  // Unicode Sign

	//Inline
	fmt.Printf("%d\t%b\t%o\t%#x\t%U\n", X, X, X, X, X)

	// Sprint ... allow us to initilize the printed String to variable
	XXX := fmt.Sprintf("%d\t%b\t%o\t%#x\t%U\n", X, X, X, X, X)
	fmt.Println(XXX)

}

// Print ... Print
// Printf ... Formated Print
// Println ... Print and Add Line at End
// FPrint ... FilePrint

func customType() {
	type Sunil int

	var Y int
	var X Sunil
	X = 47
	Y = 100
	fmt.Printf("%+v with Type %T\n", X, X)

	// Not Allowed as They are of Different Type
	// X = Y

	// Type Casting
	X = Sunil(Y)
	fmt.Println(X)

}
