package main

import "fmt"

func main() {
	// sliceLiteral()
	// slice2Literal()
	// append2Slice()
	// removeSlice()
	// makeSlice()
	multiDimensionSlice()
}
func multiDimensionSlice() {
	// list := make([][]int, 10, 11) // Type, Length, Capacity
	list := [][]int{} // Type, Length, Capacity

	B := []int{1, 2, 3}
	A := []int{7, 8, 9}

	list = append(list, A, B)
	fmt.Println(list)
}

func makeSlice() {
	list := make([]int, 10, 11) // Type, Length, Capacity
	fmt.Println(list)
	fmt.Println(len(list))
	fmt.Println(cap(list))
	list[5] = 69
	// list[11] = 12 // Not Allowed

	list = append(list, 10, 20, 30) // Allowed
	fmt.Println(list)
	fmt.Println(len(list))
	fmt.Println(cap(list))

}

func removeSlice() {
	X2 := []int{}
	for i := 0; i < 10; i++ {
		X2 = append(X2, i)
	}
	X2 = append(X2[:7], X2[9:]...) // Slice, ... is the Unwrapping of Slice
	fmt.Println(X2)
}

func append2Slice() {
	X2 := []int{}
	for i := 0; i < 100; i++ {
		X2 = append(X2, i)
	}
	X2 = append(X2, 19, 13, 23)
	fmt.Println(X2)
}

func sliceLiteral() {
	// X1:= type{} // Composite Literal
	X1 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(X1)

	for _, val := range X1 {
		fmt.Print(val, "\t")
	}
	fmt.Println()
}

func slice2Literal() {
	X2 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(X2[:])   // From Begin -> End
	fmt.Println(X2[3:5]) // 5 Index is Excluded
	fmt.Println(X2[:3])  // From Start to Index 2
	fmt.Println(X2[3:])  // From Index 3 to End

	for _, v := range X2[4:] {
		fmt.Printf("%v\t", v)
	}

	fmt.Println()
}
