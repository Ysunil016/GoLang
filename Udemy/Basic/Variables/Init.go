package main

import "fmt"

func main() {
	fmt.Println("Sunil Yadav")

	// for i := 1; i < 100; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Println(i)
	// 	}
	// }
	// base()
	fmt.Println(`Variable Decleration`)
	declareShort()
	fmt.Println("\nZero Values")
	zeroValues()
}

func base() {
	n, _ := fmt.Println(1, 2, 3, 4, "Sunil", "Yadav", true)
	fmt.Println("Total Byte Size ", n)
}

// GlobalVar ... Global Variable Decleration
var GlobalVar = 1000

// OUT ... Exposed to Outside the Package, as its Capital
var OUT = 2000

// Z ... Declaring Z with DataType Int
var Z int // Default 0

// X ... String Type of Variable
var X = "Sunil Yadav"

func declareShort() {
	x := 42 // Short Declaration Operatior
	fmt.Println(x)
	x = 50
	fmt.Println(x)
	y := 100 + 24 // Expession
	fmt.Println(y)
	var G = 400                           // Auto Type Detection (Integer)
	fmt.Printf("%v with Type %T\n", G, G) // %v is for Value and %T is for Type of the Variable
	var z int                             // Default is 0, Lazy Declaration
	fmt.Println(z)
	z = 100
	fmt.Println(z)

	Z = 140
	fmt.Println(Z)

	fmt.Println(GlobalVar)

	fmt.Printf("%T\n", X)

	// This Cannot be Done as X is Static Typed for String, No Other Type of Data Allocation is Allowed
	// X = 5
}

// zeroValues ... is the Default Values of an Variable when Its Not Initialized
func zeroValues() {
	var X int
	var Y bool
	var Z string
	var P float32

	fmt.Println(X) //0
	fmt.Println(Y) // false
	fmt.Println(Z) // ""
	fmt.Println(P) // 0
}
