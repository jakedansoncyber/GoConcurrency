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

	for i := 0; i < 10; i++ {
		// grab a random id
		id := rnd.Intn(10) + 1
		// Add the amount of tasks or go routines that are going to execute each time
		waitGroup.Add(2)
		// wrap if statement with anon func so we can call go routine
		go func(id int, waitGroup *sync.WaitGroup) {
			// query cache first; if id of book is in cache, grab it
			if book, ok := queryCache(id); ok {
				fmt.Println("from cache")
				fmt.Println(book.ToString())
			}
			// tells the waitGroup that we created that this task is done
			waitGroup.Done()
		}(id, waitGroup)

		go func(id int, waitGroup *sync.WaitGroup) {
			// query database if book is not in cache
			if book, ok := queryDatabase(id); ok {
				fmt.Println("from database")
				fmt.Println(book.ToString())
			}
			// tells the waitGroup that we created that this task is done
			waitGroup.Done()
		}(id, waitGroup)
		// If book not in cache then print (shouldn't hit this)
		//fmt.Printf("Book not found with id: '%v'", id)
		//time.Sleep(150 * time.Millisecond)
	}

	// wait for all tasks to be done
	waitGroup.Wait()
}

// returns a book & a bookean if exists
func queryCache(id int) (Book, bool) {
	book, ok := cache[id]
	return book, ok
}

func queryDatabase(id int) (Book, bool) {
	// fake time it takes to query a database
	time.Sleep(100 * time.Millisecond)

	// iterate through slice of books defined in book.go
	for _, book := range books {
		// if book id is the same as id provided, means they are the same book...
		if book.ID == id {
			// put the book in the cache
			// ## Potential Shared Memory Problem! ##
			//cache[id] = book // without mutexs go routines will try to write & access this cache at the same time (access the same memory)
			return book, true
		}
	}

	// else return an empty book & say false boolean
	return Book{}, false
}
