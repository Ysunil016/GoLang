package main

import "fmt"

func main() {
	// ex1()
	// ex2()
	// ex3()
	ex4()
}

func ex4() {
	ASS := struct {
		name  string
		val   map[string]string
		songs []string
	}{
		name:  "Sunil",
		val:   map[string]string{"name": "Sanjay", "age": "24"},
		songs: []string{"Sunil", "Yadav"},
	}
	fmt.Println(ASS.name)
	fmt.Println(ASS.val["name"])
	for _, V := range ASS.songs {
		fmt.Println(V)
	}

}

func ex3() {
	type vehicle struct {
		color string
		doors int
	}

	type sedan struct {
		vehicle
		luxury bool
	}
	type truck struct {
		vehicle
		fourWheels bool
	}

	Truck := truck{
		fourWheels: true,
		vehicle: vehicle{
			color: "Red",
			doors: 4,
		},
	}

	Sedan := sedan{
		luxury: true,
		vehicle: vehicle{
			color: "Green",
			doors: 6,
		},
	}

	fmt.Print("Color : ", Truck.color, "Doors : ", Truck.doors, "4 Wheels : ", Truck.fourWheels)
	fmt.Println()
	fmt.Print("Color : ", Sedan.color, "Doors : ", Sedan.doors, "Luxury : ", Sedan.luxury)
}

func ex2() {
	type person struct {
		fname  string
		lname  string
		icream []string
	}

	P1 := person{
		fname:  "Sunil",
		lname:  "Yadav",
		icream: []string{"Vanilla", "Chocolate"},
	}
	P2 := person{
		fname:  "Sanjay",
		lname:  "Yadav",
		icream: []string{"Scotch", "Wine"},
	}
	P3 := person{
		fname:  "Nirmala",
		lname:  "Yadav",
		icream: []string{"Scotch", "Wine"},
	}

	MAP := map[string]person{
		P1.fname: P1,
		P3.fname: P3,
		P2.fname: P2,
	}

	for k, v := range MAP {
		fmt.Println(k, v)
	}

}

func ex1() {
	type person struct {
		fname  string
		lname  string
		icream []string
	}

	P1 := person{
		fname:  "Sunil",
		lname:  "Yadav",
		icream: []string{"Vanilla", "Chocolate"},
	}
	P2 := person{
		fname:  "Sanjay",
		lname:  "Yadav",
		icream: []string{"Scotch", "Wine"},
	}
	P3 := person{
		fname:  "Nirmala",
		lname:  "Yadav",
		icream: []string{"Scotch", "Wine"},
	}

	for _, v := range P1.icream {
		fmt.Println(v)
	}
	for _, v := range P2.icream {
		fmt.Println(v)
	}
	for _, v := range P3.icream {
		fmt.Println(v)
	}
}
