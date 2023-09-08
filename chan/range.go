package main

import (
	"fmt"
	"sync"
)

/*
This code demonstrates the use of goroutines and channels in Go to perform concurrent operations. It creates a channel
for integers, launches two goroutines, one to send integers to the channel and another to receive and print them, and
finally waits for the goroutines to complete using a sync.WaitGroup.
*/
func rangeOnClosedChan() {
	intStream := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range intStream {
			fmt.Printf("Read %v\n", val)
		}
	}()

	go func(ch chan<- int) {
		defer close(ch)
		for i := 1; i <= 10; i++ {
			ch <- i
		}
	}(intStream)

	wg.Wait()
}
