package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string, 1)

	go func(ch chan<- string) { // receiving channel
		ch <- "Hello World!"
		fmt.Println("Finishing goroutine")
	}(channel)

	time.Sleep(time.Second)

	message := <-channel
	fmt.Println(message)

}

// Receive only channel
func receiving(ch <-chan string) {
	msg := <-ch
	fmt.Println(msg)
	//ch <- "Test" // Error: send to receive-only type <-chan string
}

func sending(msg string, ch chan<- string) {
	ch <- msg
	//test := <-ch // ERROR: receive from send-only type chan>- string
}
