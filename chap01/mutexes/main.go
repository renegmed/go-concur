package main

import (
	"sync"
	"time"
)

// Counter keep tracks of locks
type Counter struct {
	sync.Mutex // i.e. makes Counter requires locking
	value      int
}

func main() {
	counter := Counter{}
	for i := 0; i < 10; i++ {
		go func(i int) {
			counter.Lock()
			counter.value++ // do read value of counter and also write to counter
			defer counter.Unlock()
		}(i)
	}
	time.Sleep(time.Second)
	counter.Lock()
	defer counter.Unlock()
	println(counter.value) // will read counter value thus require locking
}
