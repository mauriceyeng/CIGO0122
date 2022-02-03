package main

import (
	"fmt"
	"sync"
)

func main() {
	channel := make(chan string)
	var wait sync.WaitGroup
	wait.Add(1)
	go func() {
		channel <- "Hello from anonymous"
		fmt.Println(1)
		wait.Done()
	}()
	message := <-channel
	fmt.Println(message)
	wait.Wait()
}
