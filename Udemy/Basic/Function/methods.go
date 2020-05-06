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

func main() {
	P1 := Agent{
		person: person{
			first: "Sunil",
			last:  "Yadav",
		},
		rtk: true,
	}

	P2 := Agent{
		person: person{
			first: "Sanjay",
			last:  "Yadav",
		},
		rtk: false,
	}

	fmt.Println(P1)
	// P1.speak()
	P2.speak()
}

func (agent Agent) speak() { // Agent Type has Access to this Method
	fmt.Println("Agent Calls ", agent.first)
}
