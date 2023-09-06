package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var sharedLock sync.Mutex

const RunFor = 1 * time.Second

func greedy() {
	defer wg.Done()
	count := 0
	for begin := time.Now(); time.Since(begin) <= RunFor; {
		sharedLock.Lock()
		time.Sleep(3 * time.Nanosecond)
		sharedLock.Unlock()
		count++
	}
	fmt.Printf("Greedy executed %v loops\n", count)
}

func yield() {
	defer wg.Done()
	count := 0
	for begin := time.Now(); time.Since(begin) <= RunFor; {
		sharedLock.Lock()
		time.Sleep(1 * time.Nanosecond)
		sharedLock.Unlock()

		sharedLock.Lock()
		time.Sleep(1 * time.Nanosecond)
		sharedLock.Unlock()

		sharedLock.Lock()
		time.Sleep(1 * time.Nanosecond)
		sharedLock.Unlock()

		count++
	}
	fmt.Printf("Yield executed %v loops\n", count)

}

func main() {
	wg.Add(2)
	go greedy()
	go yield()
	wg.Wait()
}
