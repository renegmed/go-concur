package main

import (
	"fmt"
	"strings"
	"sync"
)

var wait sync.WaitGroup

func main() {

	//toUpperSync("Hello Callbacks", func(v string) {
	//	fmt.Printf("Callback: %s\n", v)
	//})

	toUpperAsycMultiLevel()

	//assembler()
}

func toUpperSync(word string, f func(string)) {
	f(strings.ToUpper(word))
}

func toUpperAsync(word string, f func(string)) {
	go func() {
		f(strings.ToUpper(word))
	}()
}

func toUpperAsycMultiLevel() {
	wait.Add(1)
	toUpperAsync("Hello Asynch Callbacks ", func(v string) {
		// level1 would pass the string param to level2 and let level2 print the strings
		toUpperAsync(fmt.Sprintf("Caller: %s\n", v), func(v string) {
			fmt.Printf("Callback: %s\n", v)
			wait.Done()
		})

	})

	fmt.Println("Waiting async response.....")
	wait.Wait()
}

// call method 1 and wait for callback for returned result 1
// call method 2 and wait for callback for returned result 2
// call method 3 and wait for result 1 and 2 and then assemble into result 3
// call method 4 and wait for result 3, process and return result 4
// call them all methods and the same time using asynch technique

// func assembler() {
//   wait.Add(1)
//
//   fmt.Println("Waiting for final result.....")
//   wait.Wait()
// }
//
// func assemble(f1 func(string), f2 func(string), f func(string)) string {
// 	go func() {
// 		return
// 	}()
// }
