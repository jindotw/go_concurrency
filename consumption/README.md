# Go Goroutine Memory Consumption

This Go program is designed to explore the memory consumption characteristics of creating and running a large number of goroutines in a Go application. It measures the memory used per goroutine by launching a significant number of goroutines that all block on an unbuffered channel and then calculates and prints the memory consumption per goroutine in kilobytes.

## Purpose

The primary purpose of this program is to demonstrate how Go's goroutines behave in terms of memory consumption when creating and running a substantial number of them concurrently. It is useful for:

- **Educational Purposes**: This program can be used as an educational tool to understand how goroutines and memory management work in the Go programming language.

- **Performance Analysis**: Developers can use this program to analyze how their Go applications handle a large number of concurrent tasks and assess the associated memory usage.

- **Memory Profiling**: It can serve as a starting point for memory profiling and optimization of Go applications that rely heavily on goroutines.

## How It Works

1. The program defines a function called `consumed` that measures and returns the memory consumption by invoking the Go garbage collector (`runtime.GC()`) and reading memory statistics using `runtime.ReadMemStats()`.

2. It creates a large number of goroutines (specified by the `numGoroutines` constant) that all block on an unbuffered channel. Each goroutine is essentially idle, waiting for a signal.

3. Before and after creating the goroutines, the program calls the `consumed` function to measure memory consumption. The difference between these two measurements provides the memory used by the goroutines.

4. The program calculates and prints the memory consumption per goroutine in kilobytes, providing insight into how much memory is allocated for each concurrent goroutine.

## Usage

To run the program and observe memory consumption, follow these steps:

1. Make sure you have the Go programming language installed on your system.

2. Clone this repository or download the `main.go` file.

3. Open a terminal and navigate to the directory containing `main.go`.

4. Run the program using the `go run` command:

   ```
   go run main.go
   ```

5. The program will output the memory consumption per goroutine in kilobytes.
