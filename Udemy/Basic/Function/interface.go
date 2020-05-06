package main

import "fmt"

type person struct {
	first string
	last  string
}

// Agent ...
type Agent struct {
	person
	rtk bool
}

// Funk ...
type Funk struct {
	rtk bool
}

// Value Can have Multiple Types
type human interface {
	speak()
}

// Empty Interface....  Every Method Implemets this Interface
type all interface {
}

func bar(h human) {
	h.speak()
	switch h.(type) {
	case person:
		fmt.Println("Person Bar is Taking Human ", h.(person).first) // Assertion
	}

}

func (p person) speak() {
	fmt.Println("Person is Speaking ", p.first, p.last)
}

func (a Agent) speak() {
	fmt.Println("Agent is Speaking ", a.person.first, a.person.last)
}

func (f Funk) speak() {
	fmt.Println("Funk is Speaking ", f)
}

func main() {

	P1 := Agent{
		person: person{
			first: "Sunil",
			last:  "Yadav",
		},
		rtk: true,
	}

	// P2 := Agent{
	// 	person: person{
	// 		first: "Sanjay",
	// 		last:  "Yadav",
	// 	},
	// 	rtk: false,
	// }

	fuk := Funk{
		rtk: true,
	}

	// P1.speak()
	// P2.speak()
	// Per.speak()

	bar(P1)
	bar(fuk)

}
