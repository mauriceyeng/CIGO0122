package main

import "fmt"

//var class map[int]string

func main() {
	//declaring class
	student_class := map[int]string{
		1: "aditya",
		2: "maurice",
		3: "devanshu",
	}

	//another class
	student_class2 := map[int]string{
		69:  "tony",
		420: "ezeikel",
		707: "killua",
	}

	student_class[5] = "Vishnu"
	student_class[7] = "kishore"

	//deleting element
	delete(student_class, 5)

	//new map edit
	new_map := student_class
	new_map[3] = "alex"

	//printing map
	for roll_no, name := range student_class {
		fmt.Println("Roll no: ", roll_no, "Name: ", name)
	}

	//printing second map
	for roll_no, name := range student_class2 {
		fmt.Println("Roll no: ", roll_no, "Name: ", name)
	}

	//key values
	for key, value := range student_class2 {
		student_class[key] = value
	}

	//defining rows/columns
	/*
		R-row(user input)
		C->col (user input)
		arr:=make([][]int,R)

		for i:=0;i<R;i++{
			arr[i]=make([]int,C)
		}
		arr={{0,0,0,0},{0,0,0},{0,0,0,0}}
	*/
}
