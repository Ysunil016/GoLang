package main

import "fmt"

func main() {
	// preDeclareConst()
	// initStatement()
	// ifElse()
	switchC()
}

func switchC() {
	switch {
	case false:
		fmt.Println("Sunil")
	case true:
		fmt.Println("Yadav")
		fallthrough // Allow to Exe all Remaining Statements if Found True
	case true:
		fmt.Println("True")
	case false:
		fmt.Println("False")
	}

	X := 100
	switch X {
	case 10, 20, 30, 100:
		fmt.Println("10,20,30,100")
	case 11:
		fmt.Println("100")
	default:
		fmt.Println("Default")
	}
}

func ifElse() {
	X := 80
	if X == 4 {
		fmt.Println("40")
	} else if X == 80 {
		fmt.Println(80)
	} else {
		fmt.Println("Sunil")
	}
}

func preDeclareConst() {
	if true {
		fmt.Println("True")
	}
	if false {
		fmt.Println("False")
	}
	if !true {
		fmt.Println("Not True")
	}
	if !false {
		fmt.Println("Not False")
	}

	if 2 == 2 {
		fmt.Println("2 == 2")
	}
	if !(2 != 2) {
		fmt.Println("2 != 2")
	}
}

func initStatement() {
	if X := 43; X == 43 {
		fmt.Println("43")
	}
}
