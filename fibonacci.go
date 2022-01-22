package main

import "fmt"

func fibo(num int) int {
	if num == 0 || num == 1 {
		return num
	}

	if num < 0 {
		fmt.Println("invalid number!")
		return -1
	}
	result := fibo(num-1) + fibo(num-2)
	return result

}
func main() {
	fmt.Println(fibo(4))
}
