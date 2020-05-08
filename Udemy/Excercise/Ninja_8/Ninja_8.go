package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

func main() {
	// marshalUser()
	// unMarshal()
	// encodingJs()
	sortStringInt()
	// sortingPerson()
}

// User ... Struct
type User struct {
	Username string `json:"id"`
	Email    string
	City     string
}

// Person ...
type Person struct {
	Name string
	Age  int
}

type sortByAge []Person

func (a sortByAge) Len() int           { return len(a) }
func (a sortByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type sortByName []Person

func (a sortByName) Len() int           { return len(a) }
func (a sortByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

func sortStringInt() {
	Stringss := []string{"Sunil", "Yadav", "Sanjay", "Nirmala", "Priyanka", "Tripti", "Aastha", "Agastya"}
	Intss := []int{3, 2, 5, 4, 2, 4, 23, 23, 43, 45, 65, 23, 56, 7, 5, 3, 4, 65, 34, 34, 65, 2}

	sort.Strings(Stringss)
	sort.Ints(Intss)
	fmt.Println(Stringss)
	fmt.Println(Intss)
}

func sortingPerson() {
	P1 := Person{"Sunil", 24}
	P2 := Person{"Anuj", 26}
	P3 := Person{"Kamla", 36}
	P4 := Person{"Rashmi", 24}

	Staff := []Person{P1, P2, P3, P4}
	// sort.Sort(sortByAge(Staff))
	sort.Sort(sortByName(Staff))
	fmt.Println(Staff)
}

func encodingJs() {
	UserOne := User{"Sunil", "sunil016@yahoo.com", "Mumbai"}
	UserTwo := User{"Sanjay", "sunil016@yahoo.com", "Panvel"}
	UserThree := User{"Nirmala", "sunil016@yahoo.com", "Navi Mumbai"}
	Collection := []User{UserOne, UserTwo, UserThree}
	err := json.NewEncoder(os.Stdout).Encode(Collection)
	if err != nil {
		fmt.Println("Issue")
	}
}

func marshalUser() {
	UserOne := User{"Sunil", "sunil016@yahoo.com", "Mumbai"}
	UserTwo := User{"Sanjay", "sunil016@yahoo.com", "Panvel"}
	UserThree := User{"Nirmala", "sunil016@yahoo.com", "Navi Mumbai"}

	Collection := []User{UserOne, UserTwo, UserThree}
	by, err := json.Marshal(Collection)
	if err != nil {
		fmt.Println(err)
		return
	}
	os.Stdout.Write(by)
	fmt.Println()
}

func unMarshal() {
	valueString := `[{"id":"Sunil","Email":"sunil016@yahoo.com","City":"Mumbai"},{"id":"Sanjay","Email":"sunil016@yahoo.com","City":"Panvel"},{"id":"Nirmala","Email":"sunil016@yahoo.com","City":"Navi Mumbai"}]`
	stringByte := []byte(valueString)

	userList := []User{}
	err := json.Unmarshal(stringByte, &userList)
	if err != nil {
		fmt.Println(err)
	}
	for _, k := range userList {
		fmt.Println(k.Username)
	}
}
