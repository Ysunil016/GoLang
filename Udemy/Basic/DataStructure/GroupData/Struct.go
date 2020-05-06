package main

import "fmt"

func main() {
	// structU()
	// anonymusStruct()
	embeddedStruct()
}

func structU() {
	type person struct {
		fname string
		lname string
	}

	p1 := person{
		fname: "Sunil",
		lname: "Yadav",
	}

	p2 := person{
		fname: "Sanjay",
		lname: "Yadav",
	}

	fmt.Println(p1.fname)
	fmt.Println(p2)

}

func anonymusStruct() {
	P := struct {
		fname string
		lname string
	}{
		fname: "Sunil",
		lname: "Yadav",
	}

	fmt.Println(P)
}

func embeddedStruct() {
	type person struct {
		name string
		age  int
	}
	type agent struct {
		person
		active bool
	}

	Agent := agent{
		person: person{
			name: "Sunil",
			age:  24,
		},
		active: true,
	}
	fmt.Println(Agent)
	fmt.Println(Agent.person.age)
	fmt.Println(Agent.active)
}
