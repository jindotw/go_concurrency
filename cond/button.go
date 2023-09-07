package main

import (
	"fmt"
	"sync"
	"time"
)

func subscribe(cond *sync.Cond, fn func()) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		wg.Done()
		cond.L.Lock()
		cond.Wait()
		fn()
		cond.L.Unlock()
	}()
	wg.Wait()
}

func clicked() {
	eventCount := sync.WaitGroup{}
	eventCount.Add(3)
	cond := sync.NewCond(&sync.Mutex{})

	subscribe(cond, func() {
		fmt.Println("Maximizing window")
		eventCount.Done()
	})
	subscribe(cond, func() {
		fmt.Println("Playing music")
		eventCount.Done()
	})
	subscribe(cond, func() {
		fmt.Println("Mouse clicked")
		eventCount.Done()
	})

	sleep := randInt() * 25
	fmt.Printf("Gonna sleep for %d milliseconds\n", sleep)
	time.Sleep(time.Duration(sleep) * time.Millisecond)
	fmt.Println("Triggering event")
	cond.Broadcast()

	eventCount.Wait()
}
