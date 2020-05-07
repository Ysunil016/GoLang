package main

import "fmt"

func main() {
	me2()
}

func me2() {
	X := 43
	foo2(&X)
	fmt.Println(X)
}

func foo2(t *int) {
	fmt.Println(*t)
	*t = 100
}

func me1() {
	X := 43
	foo1(X)
	fmt.Println(X)
}

func foo1(t int) {
	fmt.Println(t)
	T := 100
	fmt.Println(T)
}

func init2() {
	var c int
	c = 100
	fmt.Println(c)  // Value
	fmt.Println(&c) //Address in Memory
	fmt.Printf("%+T\n", c)
	fmt.Printf("%+T\n", &c)
	fmt.Printf("%+v\n", *&c) // Value Stored at the Address

	var adV *int = &c
	var d **int = &adV
	fmt.Println(adV)
	fmt.Println(**d)

	*adV = 200
	fmt.Println(**d)
}
