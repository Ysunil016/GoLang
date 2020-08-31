package main

// PASSPHRASE is Used as Secrete Key
const (
	ConnHOST   = "localhost"
	EncPORT    = "9761"
	DcrPORT    = "9762"
	ConnTYPE   = "tcp"
	PASSPHRASE = "_#RootWesee@1234"
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
