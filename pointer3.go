package main

import "fmt"

func main() {
	name := new(string)
	*name = "golang"
	fmt.Println(*name)
}
