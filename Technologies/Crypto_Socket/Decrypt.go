package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
)

// Decryptor ...
func Decryptor() {
	fmt.Println("Starting Decryptor...")

	// listen on port 8000
	ln, _ := net.Listen(ConnTYPE, ":9762")

	// accept connection
	conn, _ := ln.Accept()

	// Making Dedicated Task
	for {
		for {
			// get message, output
			msgReceived, _ := bufio.NewReader(conn).ReadString('\n')

			byteReceived, _ := ioutil.ReadFile("app.txt")
			msgReceived = string(byteReceived)

			// Received Message for Decryption
			fmt.Print(msgReceived)
			// Dcrypting Message
			decipheredText := decrypt([]byte(msgReceived), PASSPHRASE)

			// Sending Dcrypted Message
			conn.Write([]byte(decipheredText))
		}
	}

}
