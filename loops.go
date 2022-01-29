package main

import "fmt"

/*

 */
func main() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(i)

		}
	}
	/*for true {
		fmt.Println("infinite execution")

	}*/

	//dowhile
	i := 0
	for {
		fmt.Println(i)
		i++
		if i >= 3 {
			break
		}
	}
}
