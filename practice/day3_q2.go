package main

import (
	"errors"
	"fmt"
)

//hamming distance

func Distance(a, b string) (int, error) {
	hD := 0

	if len(a) == len(b) {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				hD++
			}
		}
		return hD, nil
	}
	return hD, errors.New("DNAs are not equal")
	panic("implement the Distance function")

}

func main() {
	a := "GAGCCTACTAACGGGAT"
	b := "CATCGTAATGACGGCCT"
	dist, err := Distance(a, b)
	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("the hamming distance is: ", dist)
	}

}
