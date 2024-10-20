package main

import (
	"fmt"
	"sync"
	"time"
)

// Concurrency is the ability of a program to manage multiple tasks at the same time by interleaving their execution, while parallelism is the simultaneous execution of multiple tasks or processes, typically utilizing multiple CPU cores

var m = sync.RWMutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func Goroutines() {
	t0 := time.Now()

	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}

	wg.Wait()

	fmt.Printf("Total execution time: %v \n ", time.Since(t0))
	fmt.Printf("The results are %v \n ", results)

}

func dbCall(i int) {
	var delay float32 = 500
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the db is: ", dbData[i])
	save(dbData[i])
	log()
	wg.Done()
}

func save(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log() {
	m.RLock()
	fmt.Printf("\nThe current results are: %v", results)
	m.RUnlock()
}
