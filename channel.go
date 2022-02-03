package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string, 1) //1 is number of channels
	//var wait sync.WaitGroup
	//wait.Add(1)
	go func() {
		channel <- "Hello from anonymous"
		channel <- "10"
		time.Sleep(time.Second * 2)
		fmt.Println(1)
		//	wait.Done()
	}()
	message := <-channel
	mess := <-channel
	fmt.Println(message)
	//wait.Wait()
}
