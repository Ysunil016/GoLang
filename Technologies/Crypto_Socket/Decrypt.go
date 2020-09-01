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
			msgReceived, _ := bufio.NewReader(conn).ReadBytes('\n')

			byteReceived, _ := ioutil.ReadFile("app.txt")
			msgReceived = byteReceived

			// Dcrypting Message
			decipheredText := decrypt(msgReceived, PASSPHRASE)

			fmt.Println(decipheredText)
			// Sending Dcrypted Message
			conn.Write(decipheredText)
		}
	}

}
