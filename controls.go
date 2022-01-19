package main

import "fmt"

/*
1. if-else statement
	if(condition){
		//statements

	}
	if(condition){
		//statements
	}else{
		//statements
	}
*/
func main() {
	fmt.Print("Enter a number: ")
	var input int
	fmt.Scanln(&input)

	if input%2 == 0 {
		fmt.Println("hey you are even")

	} else {
		fmt.Println("hey you are odd")
	}

	if x := 10; x%2 == 0 { //go can support multiple condition in single if
		fmt.Println("hey you are even")

	} else {
		fmt.Println("hey you are odd")
	}
}

/*switch case
switch(data){
case:

}
*/
