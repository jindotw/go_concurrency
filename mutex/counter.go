package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
	lock  sync.RWMutex
	wg    sync.WaitGroup
}

func readCount(loop int, counter *Counter) {
	counter.lock.RLock()
	defer func() {
		counter.wg.Done()
		counter.lock.RUnlock()
	}()
	fmt.Printf("Read Loop %02d: counter value is %v\n", loop, counter.count)
}

func updateCount(loop int, counter *Counter, incrBy int) {
	counter.lock.Lock()
	defer func() {
		counter.wg.Done()
		counter.lock.Unlock()
	}()
	counter.count += incrBy
	fmt.Printf("Update Loop %02d: counter value is %v\n", loop, counter.count)
}

func runCounter() {
	counter := &Counter{
		count: 0,
		lock:  sync.RWMutex{},
		wg:    sync.WaitGroup{},
	}
	for i := 0; i < 50; i++ {
		if i%4 == 0 {
			counter.wg.Add(1)
			go func(pos int) {
				updateCount(pos, counter, 1)
			}(i)
		}
		counter.wg.Add(1)
		go func(pos int) {
			readCount(pos, counter)
		}(i)
	}
	counter.wg.Wait()
}
