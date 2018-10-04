package channelsingleton

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestStartInstance(t *testing.T) {
	singleton := GetInstance()
	singleton2 := GetInstance()

	n := 5 //5000

	for i := 0; i < n; i++ {
		// if i == 15 {
		// 	go singleton.SubtractOne()
		// 	go singleton2.SubtractOne()
		// } else {
		go singleton.AddOne()
		go singleton2.AddOne()
		// }
	}

	fmt.Printf("Before loop, current count is %d\n", singleton.GetCount())

	var val int
	for val != (n * 2) {
		log.Println("GetCount val: ", val)
		val = singleton.GetCount()
		time.Sleep(10 * time.Millisecond)
	}
	singleton.Stop()

	fmt.Printf("After loop, count is %d\n", val)
}
