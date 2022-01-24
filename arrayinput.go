package main

import "fmt"

func main() {
	myarr := [2][2]int{{1, 2}, {3, 4}}
	for index1, value1 := range myarr {
		for index2, value2 := range value1 {
			fmt.Println("element at ", index1, index2, " is ", value2)
		}
	}
}
