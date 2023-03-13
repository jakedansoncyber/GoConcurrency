// #2
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
		// run a for each
		// the sending function must close the channel when it is done though!
		for msg := range ch {
			fmt.Println(msg)
		}
		wg.Done()
	}(ch, wg)

	// send
	go func(ch chan<- int, wg *sync.WaitGroup) {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch) // must close here if: we are using the foreach & not explicitly telling how many interations in our for loop
		wg.Done()
	}(ch, wg)

	wg.Wait()
}
