package main

import (
	"bufio"
	"fmt"
	"net"
)

// Encryptor ...
func Encryptor() {
	fmt.Println("Starting Encryptor...@", EncPORT)

	// listen on port
	ln, _ := net.Listen(ConnTYPE, ConnHOST+":"+EncPORT)

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

				// ioutil.WriteFile("app.txt", cipheredText, 0777)

				// Sending the Encrypted Data
				conn.Write([]byte(cipheredText))
			}
		}()
	}

}
