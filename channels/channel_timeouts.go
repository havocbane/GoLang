package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	// ch := make(chan int) // unbuffered channel
	ch := make(chan int, 1) // using a buffered channel (1) means we don't timeout on the sender

	wg.Add(1)

	go func() {
		select {
		case ch <- 1:
			fmt.Println("Sent 1 to channel")
		case <-time.After(1 * time.Second):
			fmt.Println("Timed out sending 1 to channel")
		}
		wg.Done()
	}()

	time.Sleep(3 * time.Second)

	select {
	case result := <-ch:
		fmt.Println("Received", result, "from channel")
	case <-time.After(2 * time.Second):
		fmt.Println("Timed out receiving from channel")
	}

	wg.Wait()
}
