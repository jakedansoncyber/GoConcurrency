package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {

	// create a wait group so we can await the go routines to finish
	waitGroup := &sync.WaitGroup{}

	// create a read write mutex to fix race conditions
	mutex := &sync.RWMutex{}

	ch := make(chan Book) // channel used for cached book objects
	//dbCh := make(chan Book)    // channel used for db book objects

	// query database 10 times with each time querying a random id
	for i := 0; i < 10; i++ {
		// grab a random id
		id := rnd.Intn(10) + 1
		// Add the amount of tasks or go routines that are going to execute each time
		waitGroup.Add(2)
		// wrap if statement with anon func so we can call go routine
		go func(id int, waitGroup *sync.WaitGroup, m *sync.RWMutex, ch chan<- Book) {
			// query cache first; if id of book is in cache, grab it
			if book, ok := queryCache(id, m); ok {
				ch <- book
			} else {
				if book, ok := queryDatabase(id, m); ok {
					m.Lock()
					cache[id] = book
					m.Unlock()
					ch <- book
				}
			}
			// tells the waitGroup that we created that this task is done
			waitGroup.Done()
		}(id, waitGroup, mutex, ch)

		// create on go routine per query to handle the response
		go func(channel <-chan Book) {
			select {
			case book := <-channel:
				fmt.Println(book)
			}
			waitGroup.Done()
		}(ch)

		waitGroup.Wait()
	}

	// wait for all tasks to be done

}

// returns a book & a bookean if exists
func queryCache(id int, mutex *sync.RWMutex) (Book, bool) {
	mutex.RLock()
	book, ok := cache[id]
	mutex.RUnlock()
	return book, ok
}

func queryDatabase(id int, mutex *sync.RWMutex) (Book, bool) {
	// fake time it takes to query a database
	time.Sleep(3000 * time.Millisecond)

	// iterate through slice of books defined in book.go
	for _, book := range books {
		// if book id is the same as id provided, means they are the same book...
		if book.ID == id {
			// put the book in the cache
			mutex.Lock()
			cache[id] = book // without mutexs go routines will try to write & access this cache at the same time (access the same memory)
			mutex.Unlock()
			return book, true
		}
	}

	// else return an empty book & say false boolean
	return Book{}, false
}
