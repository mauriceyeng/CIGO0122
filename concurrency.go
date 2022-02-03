package main

import (
	"fmt"
	"sync"
	"time"
)

var wait sync.WaitGroup
var counter int
var mutex sync.Mutex

func add(hint string) {
	for i := 0; i < 3; i++ {
		mutex.Lock()
		a := counter
		a++
		time.Sleep(time.Second * 2)
		counter = a
		fmt.Println(hint, "i:=", i, " Count:=", counter)
		mutex.Unlock()
	}
	wait.Done()
}
func main() {
	wait.Add(2)
	go add("first: ")
	go add("second: ")
	wait.Wait()
	fmt.Println("final value of counter:= ", counter)

}
