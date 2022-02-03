package main

import (
	"fmt"
	"sync"
)

func main() {
	var mutex sync.Mutex
	counter := 0
	for i := 0; i < 10; i++ {
		go func() {
			mutex.Lock()
			counter++
			mutex.Unlock()

		}()
		fmt.Println(counter)
	}
	//fmt.Println(counter)
}
