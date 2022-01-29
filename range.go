package main

import "fmt"

func main() {
	nums := []int{1, 3, 2, 4, 0}
	//to self allocate index
	//for _, value := range nums
	for index, value := range nums {
		fmt.Println(index)
		fmt.Println(value)

	}
}
