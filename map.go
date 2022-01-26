package main

import "fmt"

//var class map[int]string

func main() {
	student_class := map[int]string{
		1: "aditya",
		2: "maurice",
		3: "devanshu",
	}
	student_class[5] = "Vishnu"
	student_class[7] = "kishore"

	for roll_no, name := range student_class {
		fmt.Println("Roll no: ", roll_no, "Name: ", name)
	}

}
