package main

import (
	"fmt"
	"sync"
)

var wait sync.WaitGroup
var counter int
func add(hint string){
	for i:=0;i<10;i++{
		a:=counter
		a++
		time.Sleep(time.Second*2)
		counter=a
		fmt.Println(hint,"i:=",i," Count:=",counter)
	}
	wait.Done()
}
func main() {
	wait.Add(2)
	go add("first: ")
	go add("second: ")
	wait.Wait()
	fmt.Println("final value of counter:= ",counter)