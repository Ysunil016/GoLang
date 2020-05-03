package main

import "fmt"

func main() {
	playString()
}

func playString() {
	ST := "Sunil Yadav"
	fmt.Println(ST)      // Printing String
	fmt.Println(len(ST)) // Printing Length of String
	BYTEARRAY := []byte(ST)
	fmt.Println(BYTEARRAY)        // Type Conversion to Byte Array
	fmt.Printf("%T\n", BYTEARRAY) // Type
	fmt.Printf("%b\n", BYTEARRAY) // Binary Buffer
	fmt.Printf("%X\n", BYTEARRAY) // Hexadecimal Buffer

	// Printing Byte Values of the String
	for i, v := range ST {
		fmt.Printf("At Index %d We Have Hex Value %#X\n", i, v)
	}
}
