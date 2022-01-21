package main

import "fmt"

func main() {
	var i int
	i = 10
	var j *int //declaration.. j is empty
	//j := new(int)	//declaration+assignment(j is not empty)
	//j = &i
	//*j = 100
	fmt.Println(j)
	fmt.Println(i)
}
