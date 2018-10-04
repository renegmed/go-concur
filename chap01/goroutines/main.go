package main

import "time"

func main() {
	//go helloworld()
	//time.Sleep(time.Second)

	messagePrinter := func(msg string) {
		println(msg)
	}

	// using anonymous funcion
	//go func(msg string) {
	//		println(msg)
	//} ("Hello World")

	go messagePrinter("Hello World")
	go messagePrinter("Hello goroutine")
	time.Sleep(time.Second)
}

func helloworld() {
	println("Hello World!")
}
