package main

import "sync"

func main() {
	var mutex sync.Mutex
	counter := 0
	for i := 0; i < 10; i++ {
		go func() {
			counter++
		}()
	}

}
