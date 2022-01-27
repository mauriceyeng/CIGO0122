package main

import (
	"errors"
	"fmt"
)

func main() {

	result, err := divide(4, 0)
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}

}

func divide(a, b int) (int, error) {
	if b == 0 {
		return -1, errors.New("b cant be zero")
	}
	return a / b, nil
}
