package main

import "fmt"

func main() {
	numeral()
}

func numeral() {
	S := "L"
	fmt.Println(S)

	BS := []byte(S)
	fmt.Println(BS)

	N := BS[0]

	fmt.Printf("%v and %T\n", N, N)
	fmt.Printf("%b and %T\n", N, N)
	fmt.Printf("%#X and %T\n", N, N)
}
