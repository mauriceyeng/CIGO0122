package main

import "fmt"

//function is a first class variable in GoLang

func main() {
	//number := 10
	var number int
	number = 10
	fmt.Println(number)

	//store a function as a value to a variable
	var getInt func(int) int
	getInt = func(x int) int {
		fmt.Println("In a function")
		return 10 + x
	}
	j := getInt(77)
	fmt.Println(j)

}
