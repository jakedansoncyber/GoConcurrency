package main

import (
	"fmt"
	"math/rand"
	"time"
)

var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	for i := 0; i < 10; i++ {
		// grab a random id
		id := rnd.Intn(10) + 1

		// wrap if statement with anon func so we can call go routine
		go func(id int) {
			// query cache first; if id of book is in cache, grab it
			if book, ok := queryCache(id); ok {
				fmt.Println("from cache")
				fmt.Println(book.ToString())
			}
		}(id)

		go func(id int) {
			// query database if book is not in cache
			if book, ok := queryDatabase(id); ok {
				fmt.Println("from database")
				fmt.Println(book.ToString())
			}
		}(id)
		// If book not in cache then print (shouldn't hit this)
		//fmt.Printf("Book not found with id: '%v'", id)
		time.Sleep(150 * time.Millisecond)
	}
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
			cache[id] = book
			return book, true
		}
	}

	// else return an empty book & say false boolean
	return Book{}, false
}
