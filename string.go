package main

import (
	"fmt"
	"strings"
)

func main() {
	st := "alohaaa"

	for _, value := range st {
		fmt.Println(string(value))
	}
	fmt.Println(strings.Count(st, "a"))
}
