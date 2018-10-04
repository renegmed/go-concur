package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wait sync.WaitGroup

	wait.Add(3) // number of entities

	go func() {
		time.Sleep(4 * time.Second)
		fmt.Println("Hello World 1!")
		wait.Done()
	}()

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Hello World 2!")
		wait.Done()
	}()

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Hello World 3!")
		wait.Done()
	}()
	wait.Wait()
	fmt.Println("Done waiting. Continue working")

	//f := func(id int) {
	//	fmt.Printf("ID: %d: Hello goroutine!\n", id)
	//}

	multiGoRoutines(4, func(id int) {
		fmt.Printf("ID: %d: Hello goroutine!\n", id)
	})
}

func multiGoRoutines(count int, f func(int)) {
	var wait sync.WaitGroup
	goRoutines := count
	wait.Add(goRoutines)
	for i := 0; i < goRoutines; i++ {
		go func(goRoutineID int) {
			f(goRoutineID)
			//wait.Done()
			wait.Add(-1) // done, remove 1 goroutine
		}(i)
	}
	wait.Wait()
}
