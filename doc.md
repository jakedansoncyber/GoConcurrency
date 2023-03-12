# Wait Groups

We create a wait group, which is essentially a go routine that is there to track other go routines. This way we can wait on the go routines to finish before we end the program. This way there isn't any tasks trying to complete before the main function returns.

### Create

```go
// Create a wait group object
waitGroup := &sync.WaitGroup{}
```

Before a go routine, make sure to add that go routine to the waitgroup queue:

### Add to queue
```go
waitGroup.Add(1)
```

Now inside of your go routine, we need to tell to waitGroup when the task is done. Do this by:

### Tell the waitGroup our task is done inside the go routine:

```go
waitGroup.Done()
```

Now lets put it all together in a clear picture of creation, awaiting and completion of task(s).

### Full main function

```go
// Don't worry about implementation, just know that 
// each anon function is doing some go routine task
// and that our wait group awaits it
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
            // problem: no mutex = can access same memory space
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
			// tells the waitGroup that we created 
            // that this task is done
			waitGroup.Done()
		}(id, waitGroup)
	}

	// wait for all tasks to be done
	waitGroup.Wait()
}
```
### Problems:

We still have two main problems:
- If we want to avoid go routines accessing the same memory we need to add mutexs. 

- fmt.Println() is NOT thread safe. So we must use another technique to solve this problem.