// #1

package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(2)
	// recieve
	go func(ch <-chan int, wg *sync.WaitGroup) {
		// check channel with if statement

		// if channel is not closed, print message
		if msg, ok := <-ch; ok {
			fmt.Println(msg, ok)
		} else {
			fmt.Println("The channel was closed and cannot be accessed from.")
		}

		wg.Done()
	}(ch, wg)

	// send
	go func(ch chan<- int, wg *sync.WaitGroup) {
		close(ch) // UNEXPECTED CLOSE
		wg.Done()
	}(ch, wg)

	wg.Wait()
}
