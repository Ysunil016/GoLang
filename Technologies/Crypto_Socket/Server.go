package main

// PASSPHRASE is Used as Secrete Key
const (
	ConnHOST   = "localhost"
	EncPORT    = "9761"
	DcrPORT    = "9762"
	ConnTYPE   = "tcp"
	PASSPHRASE = "SUNIL@#SUNIL$%SUNIL*&SUNIL%SUNIL"
)

func main() {
	// Staring Encryptor
	go Encryptor()
	// Staring Decryptor
	go Decryptor()

	// Waiting Infinitly
	wait := make(chan int)
	<-wait
}
