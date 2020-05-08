package main

import (
	"fmt"
	"sort"
)

func main() {
	// pre()
	customSort()
}

type person struct {
	Name string
	Age  int
}

type byAge []person

func (a byAge) Len() int           { return len(a) }
func (a byAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type byName []person

func (a byName) Len() int           { return len(a) }
func (a byName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byName) Less(i, j int) bool { return a[i].Name < a[j].Name }

func customSort() {
	P1 := person{
		"Sunil", 24,
	}
	P2 := person{
		"Sanjay", 34,
	}
	P3 := person{
		"Nirmala", 36,
	}

	col := []person{P3, P1, P2}
	// sort.Sort(byAge(col))
	sort.Sort(byName(col))
	fmt.Println(col)
}

func pre() {
	arr := []int{3, 1, 3, 5, 2, 5, 6, 1, 3, 4}
	fmt.Println(sort.IntsAreSorted(arr))
	sort.Ints(arr)
	fmt.Println(arr)
	fmt.Println(sort.IntsAreSorted(arr))

	sA := []string{"Sunil", "Nirmala", "Sanjay"}
	sort.Strings(sA)
	fmt.Println(sA)
}
