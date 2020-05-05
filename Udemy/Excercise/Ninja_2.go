package main

import "fmt"

func main() {
	// ex1()
	// ex2()
	// ex3()
	// ex4()
	// ex5()
	ex6()
}

func ex1() {
	var X1 int = 911
	fmt.Printf("%v\t%b\t%#X\t\n", X1, X1, X1)
}

func ex2() {
	a := 100
	b := 143
	c := 143
	d := 100
	e := a == b
	f := c >= d
	g := a <= d
	h := b > d
	i := c < a
	j := c != a
	k := c < a
	fmt.Println(a, b, c, d, e, f, g, h, i, j, k)
}

// AAA ...
const (
	AAA        = 1000
	BBB string = "Sunil"
)

func ex3() {
	fmt.Println(AAA)
	fmt.Println(BBB)
}

func ex4() {
	var XX int = 16
	fmt.Printf("%v\t%b\t%#X\n", XX, XX, XX)
	XX = XX << 1
	fmt.Printf("%v\t%b\t%#X\n", XX, XX, XX)
}

func ex5() {
	TY := `Here is Some Raw String Literals, "Sunil"`
	fmt.Println(TY)
}

// YearOne ...
const (
	YearOne   = iota + 2019
	YearTwo   = iota + YearOne
	YearThree = iota + YearOne
	YearFour  = iota + YearOne
)

func ex6() {

	fmt.Println(YearOne, YearTwo, YearThree, YearFour)
}
