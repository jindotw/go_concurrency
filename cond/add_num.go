package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Buffer struct {
	data     []int
	capacity int
	cond     *sync.Cond
	wg       sync.WaitGroup
}

func (b *Buffer) produce(no int, data int) {
	defer func() {
		b.wg.Add(1)
		b.cond.L.Unlock()
		b.cond.Signal()
		b.wg.Done()
	}()
	b.cond.L.Lock()
	for len(b.data) >= b.capacity {
		b.cond.Wait()
	}
	b.data = append(b.data, data)
	fmt.Printf("Producer %d produced %d\n", no, data)
}

func (b *Buffer) consume() int {
	defer func() {
		b.cond.L.Unlock()
		b.cond.Signal()
		b.wg.Done()
	}()
	b.cond.L.Lock()
	for len(b.data) == 0 {
		b.cond.Wait()
	}
	val := b.data[0]
	b.data = b.data[1:]
	return val
}

func randInt() int {
	return rand.Intn(100-10) + 10
}

func addNum() {
	buff := &Buffer{
		data:     []int{},
		capacity: 3,
		cond:     sync.NewCond(&sync.Mutex{}),
		wg:       sync.WaitGroup{},
	}
	for i := 0; i < 10; i++ {
		buff.wg.Add(1)
		go func(no int) {
			val := randInt()
			buff.produce(no, val)
			time.Sleep(time.Duration(val) * time.Millisecond)
		}(i + 1)
	}

	for i := 0; i < 20; i++ {
		go func(no int) {
			val := buff.consume()
			fmt.Printf("\tConsumer %d consumed %d\n", no, val)
		}(i + 1)
	}

	buff.wg.Wait()
	fmt.Printf("Len of data is %d\n", len(buff.data))
}
