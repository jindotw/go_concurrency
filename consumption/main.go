package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	consumed := func() uint64 {
		runtime.GC()
		sys := runtime.MemStats{}
		runtime.ReadMemStats(&sys)
		return sys.Sys
	}

	var c <-chan interface{}
	wg := sync.WaitGroup{}
	noop := func() {
		wg.Done()
		<-c
	}
	const numGoroutines = 1e5
	wg.Add(numGoroutines)
	before := consumed()
	for i := numGoroutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()
	after := consumed()
	fmt.Printf("%.3fkb\n", float64(after-before)/numGoroutines/1e3)
}
