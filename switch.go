package main

import "fmt"

func main() {
	data := 100
	switch data {
	case 10:
		data = 100
		fmt.Println(data)
	case 100:
		data = 103
		fmt.Println(data)
		fallthrough
	case 11:
		data = 104
		fmt.Println(data)
		fallthrough

	default:
		data = 6969
		fmt.Println(data)
	}
}
