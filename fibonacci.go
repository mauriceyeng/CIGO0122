package main

import "fmt"

func fibo(num int) int {
	if num == 0 {
		return 0
	}
	if num == 1 {
		return 1
	}
	if num < 3 {
		fmt.Println("invalid number!")
		return 0
	}
	result := fibo(num-2) + fibo(num-1)
	return result

}
func main() {
	fibo(10)
}
