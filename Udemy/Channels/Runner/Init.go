package main

import "fmt"
func main() {
	channel:=make(chan int) // Channel Can Hold 2 Quantity of Value, If We Try To Insert 3rd One It will be Blocked
	go func(){
		channel<-42
	}()
	fmt.Println(<-channel)
}