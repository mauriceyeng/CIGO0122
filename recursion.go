package main

import "fmt"

func main() {

}
func fact(number int) int {
	if number == 1 || number == 0 {
		return 1
	}
	if number < 0 {
		fmt.Println("invalid number")
		return 0
	}
	result := number * fact(number-1)
	return result

}
