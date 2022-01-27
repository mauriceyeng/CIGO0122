package main

import "fmt"

type SalaryCalculator interface {
}

func describe(i interface{}) {
	fmt.Printf("Type= %T,value=%v", i, i)
	var num int
	num = i.(int)

}

func main() {
	describe("Hello")
	describe(12)
	describe(true)

}
