package main

import "fmt"

func main() {
	channel := make(chan int, 2) // Channel Can Hold 2 Quantity of Value, If We Try To Insert 3rd One It will be Blocked
	channel <- 42
	channel <- 43
	fmt.Println(<-channel)
	fmt.Println(<-channel)
}
