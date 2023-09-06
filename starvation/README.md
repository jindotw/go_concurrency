# Starvation Demonstration in Go

This Go program serves as a demonstration of the concept of "starvation" in concurrent programming. Starvation occurs when one thread or goroutine is consistently given priority over others, leading to delayed or unfairly treated execution for the latter.

## Program Overview

The program includes two functions, `greedy` and `yield`, which run concurrently using goroutines and synchronization mechanisms. The main purpose is to showcase different locking behaviors and how they can affect the execution of concurrent operations.

- **`greedy` Function**: This function represents a scenario where one goroutine is less considerate of others. It acquires a shared lock and holds it for a relatively longer duration (3 nanoseconds) before releasing it. This behavior can potentially starve other goroutines waiting to acquire the same lock.

- **`yield` Function**: In contrast, the `yield` function demonstrates a more cooperative approach. It acquires the shared lock but releases it more frequently (after just 1 nanosecond) before reacquiring it. This behavior allows other goroutines to have a chance to acquire the lock, reducing the likelihood of starvation.

## Running the Program

To run the program, ensure you have Go installed on your system. Then, execute the following command in your terminal:

```shell
go run main.go
