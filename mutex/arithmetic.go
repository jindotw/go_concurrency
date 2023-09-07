package main

import (
	"fmt"
	"sync"
)

type Arithmetic struct {
	lock  sync.Mutex
	wg    sync.WaitGroup
	count int
}

func increment(obj *Arithmetic) {
	defer func() {
		obj.lock.Unlock()
		obj.wg.Done()
	}()
	obj.lock.Lock()
	obj.count++
	fmt.Printf("Count is %v\n", obj.count)
}

func decrement(obj *Arithmetic) {
	defer func() {
		obj.lock.Unlock()
		obj.wg.Done()
	}()
	obj.lock.Lock()
	obj.count--
	fmt.Printf("Count is %v\n", obj.count)
}

func arithmetic() {
	obj := &Arithmetic{
		lock:  sync.Mutex{},
		wg:    sync.WaitGroup{},
		count: 0,
	}
	for i := 0; i < 6; i++ {
		obj.wg.Add(2)
		go increment(obj)
		go decrement(obj)
	}
	obj.wg.Wait()
}
