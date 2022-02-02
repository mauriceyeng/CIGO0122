package main

import "fmt"

func SquareofSum(n int) int {
	var total int
	for i := 1; i <= n; i++ {
		total = total + i
	}
	return total * total
}
func SumofSquares(n int) int {
	var sos int
	for i := 1; i <= n; i++ {
		sos = sos + (i * i)
	}
	return sos
}

func Difference(n int) int {
	return SquareofSum(n) - SumofSquares(n)
}
func main() {
	var inp int
	fmt.Scanln(&inp)
	fmt.Println(Difference(inp))
}
