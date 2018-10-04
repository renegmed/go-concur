package channelsingleton

import "log"

var addCh = make(chan bool)
var getCountCh = make(chan chan int)
var quitCh = make(chan bool)

type singleton struct{}

var instance singleton

func GetInstance() *singleton {
	return &instance
}

func (s *singleton) GetCount() int {
	resCh := make(chan int)
	defer close(resCh)
	getCountCh <- resCh // getCountCh receives a channel resCh
	return <-resCh      // will wait for getCountCh to process. resCh contains the last count
}

func (s *singleton) AddOne() {
	addCh <- true
}

func (s *singleton) SubtractOne() {
	addCh <- false
}

func (s *singleton) Stop() {
	quitCh <- true
	close(addCh)
	close(getCountCh)
	close(quitCh)
}

func init() {
	var count int

	go func(addCh <-chan bool, getCountCh <-chan chan int,
		quitCh <-chan bool) {
		for { // to receive messages from different channels at the same time
			select {
			// case v := <-addCh: // receive message, if true add 1 to counter 'count'
			// 	log.Println("Unload <-addCh : ", v)
			// 	if v {
			// 		count++
			// 	} else {
			// 		count--
			// 	}
			case <-addCh: // receive message, if true add 1 to counter 'count'
				log.Println("Unload <-addCh : ")
				count++
			case ch := <-getCountCh:
				log.Println("Unload <-getCountCh")
				ch <- count // send cummulative count to a channel kept by channel getCountCh
			case <-quitCh:
				log.Println("Unload <-quitCh")
				break
			}
		}

	}(addCh, getCountCh, quitCh)
}
