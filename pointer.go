package main

import "fmt"

func main() {
	i := 10
	j := &i
	*j = 100
	fmt.Println(j)
	fmt.Println(i)
}
