package main

import (
	"fmt"
	// "math/rand"
	"time"
	"sync"
)

// var m = sync.Mutex{}
// There is also another kind of mutex called RWMutex which is read write mutex
var m = sync.RWMutex{}
var waitGroup = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id3", "id4"}
var results = []string{}

func main () {
	t0 := time.Now()
	for i:= 0 ; i<len(dbData); i++ {
		waitGroup.Add(1)
		// go keywork is used to run the function concurrently - think of it like async
		go dbCall(i) // Without wait group the program will exit immediately
	}
	waitGroup.Wait()
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
	fmt.Printf("\nThe results are: %v", results)
}

func dbCall (i int) {
	// var delay float32 = rand.Float32()*2000
	var delay float32 = 2000 // Removing randomness to depict what happen when multiple modification happens
	time.Sleep(time.Duration(delay)*time.Millisecond)
	fmt.Println("The result from the database is: ", dbData[i])
	save(dbData[i])
	log()
	waitGroup.Done()
}

func save (result string){
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log() {
	m.RLock()
	fmt.Println("The result from the database is: ", results)
	m.RUnlock()
}