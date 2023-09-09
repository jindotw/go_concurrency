package main

import "fmt"

/*
This code demonstrates the use of a Go select statement to uniformly select and count occurrences of closed channels
ch1 and ch2. It creates two closed channels, then repeatedly selects from them in a loop, incrementing counters for
each channel to determine how many times each channel is selected. The output at the end displays the counts for both
channels.
*/
func uniformSelect() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	close(ch1)
	close(ch2)

	var count1, count2 int
	for i := 1; i <= 1000; i++ {
		select {
		case <-ch1:
			count1++
		case <-ch2:
			count2++
		}
	}

	fmt.Printf("count1: %v / count2: %v\n", count1, count2)
}
