package main

import "fmt"

func main() {
	// forLoop()
	// nestedLoop()
	// forInf()
	// forBreak()
	// forRange()

	// beakContinue()
	asciiPrint()
}

func asciiPrint() {
	for i := 65; i < 123; i++ {
		fmt.Printf("%c\t%#X\t%#U\n", i, i, i)
	}
}

func beakContinue() {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			continue
		} else {
			fmt.Println(i)
		}
	}
}

func forLoop() {
	for i := 1; i < 100; i++ {
		fmt.Println(i)
	}
}

func nestedLoop() {
	for i := 1; i < 100; i++ {
		for j := 1; j < 100; j++ {
			fmt.Println(i, j)
		}
	}
}

func forInf() {
	a := 0
	for a < 10 {
		fmt.Println(a)
		a++
	}
}

func forBreak() {
	a := 1
	for {
		if a > 10 {
			break
		} else {
			fmt.Println(a)
		}
		a++
	}
}

func forRange() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for _, v := range pow {
		fmt.Println(v)
	}
}
