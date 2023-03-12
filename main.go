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

	// query database 10 times with each time querying a random id
	for i := 0; i < 10; i++ {
		// grab a random id
		id := rnd.Intn(10) + 1
		// Add the amount of tasks or go routines that are going to execute each time
		waitGroup.Add(2)
		// wrap if statement with anon func so we can call go routine
		go func(id int, waitGroup *sync.WaitGroup, m *sync.RWMutex) {
			// query cache first; if id of book is in cache, grab it
			if book, ok := queryCache(id, mutex); ok {
				fmt.Println("from cache")
				fmt.Println(book.ToString())
			}
			// tells the waitGroup that we created that this task is done
			waitGroup.Done()
		}(id, waitGroup, mutex)

		go func(id int, waitGroup *sync.WaitGroup, m *sync.RWMutex) {
			// query database if book is not in cache
			if book, ok := queryDatabase(id, mutex); ok {
				fmt.Println("from database")
				fmt.Println(book.ToString())
			}
			// tells the waitGroup that we created that this task is done
			waitGroup.Done()
		}(id, waitGroup, mutex)
		// If book not in cache then print (shouldn't hit this)
		//fmt.Printf("Book not found with id: '%v'", id)
	}

	// wait for all tasks to be done
	waitGroup.Wait()
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
	time.Sleep(100 * time.Millisecond)

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
