package main

import (
	"fmt"
)

func main() {
	channel := make(chan string)
	//var waitGroup sync.WaitGroup
	//waitGroup.Add(1)

	go func() {
		//time.Sleep(time.Second)
		channel <- "First Hello World!"
		fmt.Println("Finishing goroutine 1")
		//waitGroup.Add(-1)
	}()

	go func() {
		//time.Sleep(time.Second)
		channel <- "Second Hello World!"
		fmt.Println("Finishing goroutine 2")
		//waitGroup.Add(-1)
	}()

	go func() {
		//time.Sleep(time.Second)
		channel <- "Third Hello World!"
		fmt.Println("Finishing goroutine 3")
		//waitGroup.Add(-1)
	}()
	//time.Sleep(time.Second)

	message := <-channel
	fmt.Printf("First message received from channel: %s\n", message)

	message = <-channel
	fmt.Printf("Second message received from channel: %s\n", message)

	message = <-channel
	fmt.Printf("Third message received from channel: %s\n", message)

	// this will cause blockage if this empty-channel call is not the last call
	//message = <-channel
	//fmt.Printf("Fourth message received from channel: %s\n", message)

	//waitGroup.Wait()
	//fmt.Println("Done waiting")
}
