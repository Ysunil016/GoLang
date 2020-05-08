package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Person ... Needs to be Able to Export from Base, with JSON Tags
type Person struct {
	Name    string   `json:"Fname"`
	Age     int      `json:"Age"`
	Friends []string `json:"FriendList"`
}

func main() {

	// Marsheling

	Sunil := Person{
		Name:    "Sunil",
		Age:     24,
		Friends: []string{"Rashmi", "Sanjay", "Nirmala"},
	}
	b, err := json.Marshal(Sunil)
	if err != nil {
		fmt.Println(b)
	}
	// fmt.Println(string(b))
	os.Stdout.Write(b)
	fmt.Println()

	// Unmarshaling
	var uPerson Person
	json.Unmarshal(b, &uPerson) // It takes []byte and an Struct Address
	fmt.Println(uPerson)
}
