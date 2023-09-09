package main

import (
	"fmt"
	"math/rand"
)

/*
This code demonstrates the concept of channel ownership in Go, where the channel creation, writing, and closing
responsibilities are all within the same function. It defines a function chanOwner that creates a buffered channel,
populates it with random integers, and then closes the channel. The main part of the code calls chanOwner to obtain
a channel and reads values from it, illustrating the ownership of the channel's lifecycle.
*/
func ownership() {
	chanOwner := func() <-chan int {
		intStream := make(chan int, 5)
		go func() {
			defer close(intStream)
			for i := 1; i <= 10; i++ {
				intStream <- rand.Intn(100-98) + 98
			}
		}()

		return intStream
	}

	readStream := chanOwner()
	for val := range readStream {
		fmt.Printf("I have received %d\n", val)
	}
	fmt.Println("Done receiving")
}
