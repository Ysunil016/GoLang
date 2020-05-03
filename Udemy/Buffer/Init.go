package main

import (
	"bytes"
	"fmt"
)

func main() {
	var b bytes.Buffer // A Buffer needs no initialization.
	b.Write([]byte(""))
	fmt.Println(b)
}
