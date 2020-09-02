package main

import (
	"bufio"
	"fmt"
	"net"
)

// Decryptor ...
func Decryptor() {
	fmt.Println("Starting Decryptor...@", DcrPORT)

	// listen on port
	ln, _ := net.Listen(ConnTYPE, ConnHOST+":"+DcrPORT)

	// accept connection
	conn, _ := ln.Accept()

	// Making Dedicated Task
	for {
		for {
			// get message, output
			msgReceived, _ := bufio.NewReader(conn).ReadBytes('\n')

			// byteReceived, _ := ioutil.ReadFile("app.txt")
			// msgReceived = byteReceived

			// Dcrypting Message
			decipheredText := decrypt(msgReceived, PASSPHRASE)

			fmt.Println(decipheredText)
			// Sending Dcrypted Message
			conn.Write(decipheredText)
		}
	}

}
