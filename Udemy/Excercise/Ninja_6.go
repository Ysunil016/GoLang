package main

import (
	"fmt"
	"math"
)

func main() {
	// ex1()
	// ex2()
	// ex3()
	// ex4()
	// ex5()
	// ex6()
	// ex7()
	// ex8()
	// ex9()
	// ex10()
	// ex11()
	ex12()
}

func ex12() {
	f := incrementor()
	g := incrementor()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(f())
	fmt.Println(f())
}

func incrementor() func() int {
	X := 1
	return func() int {
		X++
		return X
	}
}

func ex10() {
	XXX := 10
	{
		XXX := 20
		fmt.Println(XXX)
	}
	fmt.Println(XXX)
}

func sum(X ...int) int {
	Sum := 0
	for _, v := range X {
		Sum += v
	}
	return Sum
}

func ex9() {
	XXX := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println(exx9(sum, XXX...))
}

func exx9(f func(...int) int, nums ...int) int {
	SUNIL := []int{}
	for _, v := range nums {
		if v%2 == 0 {
			SUNIL = append(SUNIL, v)
		}
	}
	return f(SUNIL...)
}

func ex8() {
	XX := f8()
	fmt.Println(XX())
}

func f8() func() int {
	X := func() int {
		return 6969
	}
	return X
}

func ex7() {
	func() int {
		for i := 0; i < 20; i++ {
			fmt.Println(i)
		}
		return 6969
	}()
}

func ex6() {
	XXX := func() int {
		for i := 0; i < 20; i++ {
			fmt.Println(i)
		}
		return 6969
	}
	fmt.Println(XXX())
}

type square struct {
	side float64
}
type circle struct {
	radius float64
}

func ex5() {
	X := circle{
		radius: 10,
	}
	Y := square{
		side: 10,
	}
	info(X)
	info(Y)
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Print("Area of Shapeis ", s.area(), "\n")
}

func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func ex4() {
	P1 := person{
		first: "Sunil",
		last:  "Yadav",
		age:   24,
	}
	P1.speak()
}

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("Person Speaking ", p.first)
}

func ex3() {
	defer defer3()
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(foo3(arr...))
}

func foo3(val ...int) int {
	bar3()
	Sum := 0
	for _, v := range val {
		Sum += v
	}
	return Sum
}

func bar3() {
	defer func() {
		fmt.Println("DEFER BAR 3")
	}()
	fmt.Println("BAR 3")
}

func defer3() {
	fmt.Println("Defering")
}

func ex2() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(foo2(arr...))
	fmt.Println(bar2(arr))
}

func foo2(val ...int) int {
	Sum := 0
	for _, v := range val {
		Sum += v
	}
	return Sum
}

func bar2(val []int) int {
	Sum := 0
	for _, v := range val {
		Sum += v
	}
	return Sum
}

func ex1() {
	fmt.Println(foo1())
	fmt.Println(bar1())
}

func foo1() int {
	return 69
}

func bar1() (int, string) {
	return 100, "Sunil"
}
