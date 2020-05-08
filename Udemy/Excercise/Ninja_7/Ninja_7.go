package main

import "fmt"

func main() {
	// ex1()
	ex2()
}

type person struct {
	name string
	age  int
}

func changeMe(per *person) {
	// *per = person{"Sanjay", 34} // Changing Value at the Address
	per.name = "Miss Penny" // Changing Value at the Address
}

func ex2() {
	Per := person{"Sunil", 24}
	fmt.Println(Per)
	changeMe(&Per)
	fmt.Println(Per)
}

func ex1() {
	X := 40
	fmt.Println(X)
	fmt.Println(&X)
}
