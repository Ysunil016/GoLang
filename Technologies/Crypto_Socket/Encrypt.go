package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
)

// Encryptor ...
func Encryptor() {
	fmt.Println("Starting Encryptor...")

	// listen on port 8000
	ln, _ := net.Listen(ConnTYPE, ":9761")

	// Making Dedicated Task
	for {
		// Accept connection
		conn, _ := ln.Accept()

		// Handing Received Message
		go func() {
			for {
				// get message, output
				msgReceived, _ := bufio.NewReader(conn).ReadBytes('\n')
				// Received Message for Encryption
				fmt.Print(msgReceived)
				// Encrypting Message
				cipheredText := encrypt(msgReceived, PASSPHRASE)

				ioutil.WriteFile("app.txt", cipheredText, 0777)

				// Sending the Encrypted Data
				conn.Write([]byte(cipheredText))
			}
		}()
	}

}
