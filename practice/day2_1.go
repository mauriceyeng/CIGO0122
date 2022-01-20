package main

import (
	"errors"
	"fmt"
)

func factorial(number int) (fact int, err error) {
	fact = 1
	if number < 0 {
		err = errors.New("number has to be positive")
		return
	} else if number == 0 {
		return
	}
	for number > 0 {
		fact *= number
		number--
	}
	return
}
func main() {
	var num int
	fmt.Println("enter number: ")
	fmt.Scanln(&num)
	result, err := factorial(num)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("the factorial is: ", result)
	}

}
