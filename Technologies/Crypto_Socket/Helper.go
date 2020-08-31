package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"
	"log"
)

// Creating Hash
func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

// Encrypt Byte String with Password
func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Println(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)

	// // Testing Purposes
	// err = ioutil.WriteFile("app.txt", ciphertext, 0777)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	return ciphertext
}

// Decrypting Byte String with Password
func decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Println(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		log.Println(err.Error())
	}
	return plaintext
}

// func encryptFile(filename string, data []byte, passphrase string) {
// 	f, _ := os.Create(filename)
// 	defer f.Close()
// 	f.Write(encrypt(data, passphrase))
// }

// func decryptFile(filename string, passphrase string) []byte {
// 	data, _ := ioutil.ReadFile(filename)
// 	return decrypt(data, passphrase)
// }

// // Main Function
// func main() {
// 	ciphertext := encrypt([]byte("Sunil Yadav I am Living in Delhi and Ready to Join ThoughtWorks"), PASSPHRASE)
// 	fmt.Printf("Encrypted: %x\n", ciphertext)
// 	plaintext := decrypt(ciphertext, PASSPHRASE)
// 	fmt.Printf("Decrypted: %s\n", plaintext)

// 	// encryptFile("sample.txt", []byte("Hello World"), "password1")
// 	// fmt.Println(string(decryptFile("sample.txt", "password1")))
// }
