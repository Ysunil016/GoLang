package main

import "fmt"

func main() {
	map2()
}

func map2() {
	M := map[string]int{"Name": 69, "Surname": 2}
	fmt.Println(M)
	fmt.Println(M["Name"])   // Found Key
	fmt.Println(M["SUNSIS"]) // No Entry --- Zero Value

	M["Sunil"] = 100 // Adding Data to Map
	M["Sunil"] = 500 // Updating Data to Map

	fmt.Println(M)
	delete(M, "Name") // Deleting Data from Map
	fmt.Println(M)

	v, ok := M["SUNSIS"]
	fmt.Println(v, ok)

	for k, v := range M {
		fmt.Println(k, v)
	}
}
