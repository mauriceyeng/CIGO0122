package main

import "fmt"

type SalaryCalculator interface {
}

func describe(i interface{}) {
	fmt.Printf("Type= %T,value=%v\n", i, i)
	switch i.(type) {
	case string:
		fmt.Println("heyy i am string and i will do")
	case bool:
		fmt.Println("heyy i am bool and i will do")
	case int:
		fmt.Println("heyy i am int and i will do")
	default:
		fmt.Println("unsupported")
	}

}

func main() {
	describe("Hello")
	describe(12)
	describe(true)
	describe(2.1)

}
