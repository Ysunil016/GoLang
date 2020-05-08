package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hola")
	fmt.Fprint(os.Stdout, "Ola\n")
	io.WriteString(os.Stdout, "Oli\n")
}
