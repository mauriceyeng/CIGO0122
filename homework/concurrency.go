package main

import (
	"fmt"
	"time"
)

func main() {
	//var mutex sync.Mutex
	counter := 0
	for i := 1; i <= 10; i++ {
		go func(i int) {
			//mutex.Lock()
			counter++
			fmt.Println(i, "=>", counter)
			//mutex.Unlock()

		}(i)

	}
	//fmt.Println(counter)
	time.Sleep(time.Second * 3)
}
