// Select statement handles more than one channel input
// chooses between two or more channels
// executed just once
// use for loop if to execute different channels

package main

import (
	"fmt"
	"time"
)

func main() {
	helloCh := make(chan string, 1)
	goodbyeCh := make(chan string, 1)
	quitCh := make(chan bool)

	go receiver(helloCh, goodbyeCh, quitCh)
	go sendString(helloCh, "Hello!")
	//time.Sleep(time.Second)
	go sendString(goodbyeCh, "Goodbye!")
	<-quitCh // wait for other channels to finish and receive the quit message
	//fmt.Println(<-quitCh)
}

func sendString(ch chan<- string, s string) {
	ch <- s
}

func receiver(helloCh, goodbyeCh <-chan string, quitCh chan<- bool) {
	for {
		select {
		case msg := <-helloCh:
			fmt.Println(msg)
		case msg := <-goodbyeCh:
			fmt.Println(msg)
		case time := <-time.After(time.Second * 2): // time.After returns <-chan current time
			fmt.Printf("Nothing received in 2 seconds. Exiting. Current time %v\n", time)
			quitCh <- true
			break
		}
	}
}
