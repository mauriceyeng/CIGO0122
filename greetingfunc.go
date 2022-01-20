package main

import (
	"errors"
	"fmt"
)

func greetings(name string) (string, error) {
	if name == "" {
		return name, errors.New(" name cannot be empty!")
	}
	return "hello," + name, nil
}

func main() {
	//fmt.Println(greetings("Maurice"))
	greeting, err := greetings("Maurice Yeng")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(greeting)
	}
}
