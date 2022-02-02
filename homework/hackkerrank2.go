package main

import (
	"fmt"
	"strings"
)

func RunLengthEncode(input string) string {
	var count int
	a := strings.Count(input)
	for i := 0; i < a; i++ {

	}

}

func RunLengthDecode(input string) string {

}
func main() {
	var alphas string
	fmt.Scanln(&alphas)
	fmt.Println(RunLengthEncode(alphas))
	fmt.Println(RunLengthDecode(alphas))

}
