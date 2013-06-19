package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// This is a Go version of the famous stupid sleep sort algorithm.
// To sort n numbers, we run n goroutines.
// Goroutine i sleeps a[i] seconds and then sends a[i] to a channel.
// In this way the numbers come to the channel in increasing order.
func main() {
	c := make(chan int)
	for _, arg := range os.Args[1:] {
		go func(s string) {
			d, _ := strconv.Atoi(s)
			time.Sleep(time.Duration(d) * time.Second)
			c <- d
		}(arg)
	}
	for _ = range os.Args[1:] {
		fmt.Print(" ", <-c)
	}
	fmt.Println()
}
