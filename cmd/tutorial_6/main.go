package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex = sync.RWMutex{}
var dbData = [5]string{"id1", "id2", "id3", "id4", "id5"}
var wg = sync.WaitGroup{}
var results = []string{}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i, &results)
	}
	wg.Wait()
	fmt.Println("Total execution time is: ", time.Since(t0))
	fmt.Println("The results are: ", results)
}

func dbCall(i int, results *[]string) {
	delay := 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is: ", dbData[i])
	save(dbData[i], results)
	log()
	wg.Done()
}

func save(result string, results *[]string) {
	mutex.Lock()
	*results = append(*results, result)
	mutex.Unlock()
}

func log() {
	mutex.RLock()
	fmt.Printf("\n The current results are: %v", results)
	mutex.RUnlock()
}
