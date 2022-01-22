package main

import "fmt"

func rec(num int) {
	if num == 0 {
		return
	}
	if num%2 == 0 {
		fmt.Println(num + 1)
		rec(num)
	} else {
		fmt.Println(num + 2)
		rec(num)
	}
}
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
	fmt.Println(rec(3))
}
