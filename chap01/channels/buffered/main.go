package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string, 1)
	go func() {
		channel <- "First Hello World!" // after receiving message channel will be blocked
		fmt.Println("-- sent to channel first message")
		channel <- "Second Hello World!"
		fmt.Println("-- sent to channel second message")
		channel <- "Third Hello World!"
		fmt.Println("-- sent to channel third message")
		fmt.Println("Finishing goroutine")
	}()

	// go func() {
	// 	channel <- "Second Hello World!"
	// 	fmt.Println("Finishing goroutine 2")
	// }()

	// need to delay for display messages
	time.Sleep(2 * time.Second)

	message := <-channel // after unloadinf first message, channel will receive second message inside goroutine
	fmt.Println(message)

	message = <-channel // after unloadinf second message, channel will receive third message inside goroutine
	fmt.Println(message)

	message = <-channel
	fmt.Println(message)

	//channel is empty
	// message = <-channel
	// fmt.Println(message)

	// for message := range channel {
	// 	fmt.Println(message)
	// }
}
