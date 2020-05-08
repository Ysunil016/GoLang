package main

import "fmt"

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
}

func ex10() {
	lastName := map[string][]string{"Sunil": {"A", "D", "C", "B"}, "Nirmala": {"N", "I", "M", "M"}, "Sanjay": {"A", "D", "C", "B"}}
	// Random Order of Access
	fmt.Println(lastName)
	lastName["Rashmi"] = []string{"R", "A", "S", "H", "M", "I"}
	//Checking Before Deleting
	if _, ok := lastName["Sunil"]; ok {
		delete(lastName, "Sunil")
	}
	for k, v := range lastName {
		fmt.Println(k, " has ", v)
	}
}

func ex9() {
	lastName := map[string][]string{"Sunil": {"A", "D", "C", "B"}, "Nirmala": {"N", "I", "M", "M"}, "Sanjay": {"A", "D", "C", "B"}}
	// Random Order of Access
	fmt.Println(lastName)
	lastName["Rashmi"] = []string{"R", "A", "S", "H", "M", "I"}
	for k, v := range lastName {
		fmt.Println(k, " has ", v)
	}
}

func ex8() {
	lastName := map[string][]string{"Sunil": {"A", "D", "C", "B"}, "Nirmala": {"N", "I", "M", "M"}, "Sanjay": {"A", "D", "C", "B"}}
	// Random Order of Access
	fmt.Println(lastName)
	for k, v := range lastName {
		fmt.Println(k, " has ", v)
	}
}

func ex7() {
	var X [][]string
	Y := []string{"I", "J", "K"}
	Z := []string{"L", "M", "N"}
	X = append(X, Y, Z)
	fmt.Println(X)

	for i := 0; i < len(X); i++ {
		for j := 0; j < len(X[i]); j++ {
			fmt.Print(X[i][j], "\t")
		}
		fmt.Println()
	}

}

func ex6() {
	X := make([]string, 50, 50)
	X = []string{`Alabama`, `Alaska`, `Arizona`, `Arkansas`, `California`, `Colorado`, `Connecticut`, `Delaware`, `Florida`, `Georgia`, `Hawaii`, `Idaho`, `Illinois`, `Indiana`, `Iowa`, `Kansas`, `Kentucky`, `Louisiana`, `Maine`, `Maryland`, `Massachusetts`, `Michigan`, `Minnesota`, `Mississippi`, `Missouri`, `Montana`, `Nebraska`, `Nevada`, `New Hampshire`, `New Jersey`, `New Mexico`, `New York`, `North Carolina`, `North Dakota`, `Ohio`, `Oklahoma`, `Oregon`, `Pennsylvania`, `Rhode Island`, `South Carolina`, `South Dakota`, `Tennessee`, `Texas`, `Utah`, `Vermont`, `Virginia`, `Washington`, `West Virginia`, `Wisconsin`, `Wyoming`}
	fmt.Println(len(X))
	fmt.Println(cap(X))
	// fmt.Println(X)
	for i := 0; i < len(X); i++ {
		fmt.Printf("%v has Value %v\n", i, X[i])
	}
}

func ex5() {
	X := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	X = append(X[:3], X[6:]...)
	fmt.Println(X)
}
func ex4() {
	X := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	X = append(X, 52)
	fmt.Println(X)
	X = append(X, 53, 54, 55)
	fmt.Println(X)
	Y := []int{56, 57, 58, 59, 60}
	X = append(X, Y...)
	fmt.Println(X)
}

func ex3() {
	X1 := []int{41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(X1[:])
	fmt.Println(X1[1:6])
	fmt.Println(X1[7:])
	fmt.Println(X1[2:7])
}

func ex2() {
	var X1 = []int{10, 20, 30, 40}
	X1 = append(X1, 50)
	fmt.Println(X1)
	fmt.Println(len(X1))
	fmt.Println(cap(X1))
	fmt.Println(X1)
}

func ex1() {
	var X1 [5]int
	X1[2] = 30
	X1[3] = 40
	X1[4] = 50
	fmt.Printf("%+v and %T\n", X1, X1)
	for _, v := range X1 {
		fmt.Print(v, " ")
	}
}
