package main

import "fmt"

func main() {
	var slice []int

	for i := 0; i < 6; i++ {
		slice = append(slice, 20)
	}

	fmt.Println(slice)
}

//make function
/*make([]data_type, length,capacity)
for i:=0;i<6;i++{
	s= append(s,20)
}
s=append(s,25)
*/
