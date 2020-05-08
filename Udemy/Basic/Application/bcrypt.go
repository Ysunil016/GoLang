package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	password := "password123"
	encPass, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(encPass))

	// Compare Hash

	password = "password123"
	err = bcrypt.CompareHashAndPassword(encPass, []byte(password))
	if err == nil {
		fmt.Println("Password Matching")
		return
	}
	fmt.Println("Password Doesn't Match")
}
