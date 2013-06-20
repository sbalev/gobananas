package main

import (
	"fmt"
	"math/rand"
)

const (
	N = 1000
	K = 5
)

var (
	input           = make([]chan int, K)
	result          = make(chan int)
	mutex           = make(chan bool, 1)
	waitingTasks    = 0
	waitingMessages = N
)

func startWaiting(balance int) {
	mutex <- true
	waitingTasks++
	waitingMessages += balance
	if waitingTasks == K && waitingMessages == 0 {
		for i := 0; i < K; i++ {
			close(input[i])
		}
	}
	<-mutex
}

func stopWaiting(balance int) {
	mutex <- true
	waitingTasks--
	waitingMessages -= balance
	<-mutex
}

func task(i int) {
	sum := 0
	balance := 0
	for {
		startWaiting(balance)
		x, ok := <-input[i]
		if !ok {
			break
		}
		stopWaiting(balance)
		balance--
		if r := x % K; r == i {
			sum += x
		} else {
			input[r] <- x
			balance++
		}
	}
	result <- sum
}

/*
This example shows how to stop a group of communicating goroutines.
We have the following situation : K collaborating goroutines, each
one with its own input channel. Each goroutine reads data from its
channel, does some work and sends data to others. The process has
to stop when there is no more work and no unread messages.

To detect this situation, I use two variables counting the number
of waiting goroutines and the number of unreceived messages.
*/

func main() {
	for i := 0; i < K; i++ {
		input[i] = make(chan int, N)
	}
	sum := 0
	for j := 0; j < N; j++ {
		r := rand.Intn(2 * N)
		sum += r
		input[j%K] <- r
	}
	fmt.Println("sum =", sum)

	for i := 0; i < K; i++ {
		go task(i)
	}
	concurrentSum := 0
	for i := 0; i < K; i++ {
		concurrentSum += <-result
	}
	fmt.Println("concurrent sum =", concurrentSum)
}
