// #2
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
	// create channel & can have one message sitting in the channel
	ch := make(chan int, 1)

	// add 2 tasks to wait on
	wg.Add(2)
	// recieve from channel
	go func(ch chan int, wg *sync.WaitGroup) {
		// printing a message from the channel
		// <- indicates we are recieving the message from the channe;
		fmt.Println(<-ch) // prints 42
		fmt.Println(<-ch) // prints 27
		wg.Done()
	}(ch, wg)

	// send to channel
	go func(ch chan int, wg *sync.WaitGroup) {
		// put the number 42 into the channel
		ch <- 42
		ch <- 27 // this fills the one space in our unbuffered channel
		wg.Done()
	}(ch, wg)
	wg.Wait()
}
