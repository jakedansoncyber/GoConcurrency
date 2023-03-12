// Play around with concurrency before we implement it
// inside of our actual application

package main

import (
	"fmt"
	"sync"
)

func main() {
	// create waitgroup
	wg := &sync.WaitGroup{}
	// create channel
	ch := make(chan int)

	// add 2 tasks to wait on
	wg.Add(2)
	// recieve from channel
	go func(ch chan int, wg *sync.WaitGroup) {
		// printing a message from the channel
		// <- indicates we are recieving the message from the channe;
		fmt.Println(<-ch)
		wg.Done()
	}(ch, wg)

	// send to channel
	go func(ch chan int, wg *sync.WaitGroup) {
		// put the number 42 into the channel
		ch <- 42
		wg.Done()
	}(ch, wg)
	wg.Wait()
}
