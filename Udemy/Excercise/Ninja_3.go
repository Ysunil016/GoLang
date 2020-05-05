package main

import "fmt"

func main() {
	// ex1()
	// ex2()
	// ex3()
	// ex4()
	// ex5()
	// ex6()
	// ex7()
	ex8()
}

func ex8() {
	fmt.Println(true && true)
	fmt.Println(false && true)
	fmt.Println(true || true)
	fmt.Println(false && true)
	fmt.Println(!false)
}

func ex7() {
	favSport := "India"
	switch favSport {
	default:
		fmt.Println("No Pref")
	case "Cricket", "India":
		fmt.Println("You Like Cricket")
	case "Football":
		fmt.Println("You Like Football")
	}
}

func ex6() {
	switch {
	default:
		fmt.Println("No Genuine Input")
	case 200 == 100:
		fmt.Println("200==100")
	case 100 == 100:
		fmt.Println("100==100")
	case 300 == 300:
		fmt.Println("300==300")
	}
}

func ex5() {
	if true {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
	if !true {
		fmt.Println("False")
	} else {
		fmt.Println("True")
	}

	if 100 > 101 {
		fmt.Println("100 < 101")
	} else if 100 <= 100 {
		fmt.Println("100 >= 101")
	}

}

func ex4() {
	for i := 10; i < 100; i++ {
		if i%4 == 0 {
			fmt.Println(i)
		}
	}
}

func ex3() {
	a := 99
	for {
		if a < 50 {
			break
		}
		a = a - 1
		// fmt.Println(a)
	}

	for a != 20 {
		a = a - 1
		fmt.Println(a)
	}
}

func ex2() {
	for i := 32; i < 123; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}

func ex1() {
	for i := 1; i < 10000; i++ {
		fmt.Print(i, " ")
	}
}
