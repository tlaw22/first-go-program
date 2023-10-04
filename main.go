package main

import (
	"fmt"
	"time"
)

// all apps require the main function
func main() {

	// a basic for loop
	for i := 0; i < 10; i++ {
		fmt.Println("Hello, World")
	}

	// variable data types
	var xs int
	xs = 66
	var whatToSay string
	whatToSay = "Goodbye, cruel world"
	fmt.Println(whatToSay)
	fmt.Println(xs)
	// decisions
	var x int = 9
	if x > 10 {
		fmt.Println("x is greater than 10")
	} else {
		fmt.Println("x is not greater than 10")
	}
	// select case
	ch1 := make(chan string)
	ch2 := make(chan string)
	// Declare some anon funcations
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Hello from channel 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Hello from channel 2"
	}()

	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case msg := <-ch2:
		fmt.Println(msg)
	default:
		fmt.Println("No messages received")
	}

	whatWasSaud := saySomething()
	fmt.Println("The function returned", whatWasSaud)
}
func saySomething() string {
	return "ALMOST NOTHING"
}
