package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
This code demonstrates the synchronization of multiple goroutines using a channel in Go. It creates a channel of empty
structs, launches five goroutines, each of which waits to receive a signal from the channel before printing a message.
After a random sleep period, it closes the channel and waits for all the goroutines to finish using a sync.WaitGroup.
*/
func unblock() {
	wg := sync.WaitGroup{}
	stream := make(chan struct{})

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(no int, ch <-chan struct{}) {
			defer wg.Done()
			<-ch
			fmt.Printf("goroutine %v has begun\n", no)
		}(i, stream)
	}

	sleep := rand.Intn(1000-500) + 500
	fmt.Printf("Gonna sleep for %v milli\n", sleep)
	time.Sleep(time.Duration(sleep) * time.Millisecond)
	close(stream)
	wg.Wait()
}
