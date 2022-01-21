package main

import "fmt"

/*func function_name(){
	statement 1
	statement 2
}
*/

func main() {
	//a()
	/*result, re := a()
	fmt.Println(result)
	fmt.Println(re)*/
	//c()
	result, re := c()
	fmt.Println(result)
	fmt.Println(re)
	b(1, 2, 3, 4, 5, 6, 7)

}

func a() (int, string) {
	//fmt.Println(12)
	return 122, "beluga"
}

func b(args ...int) {
	for _, v := range args {
		fmt.Print(v)
	}
}

//declare values in initialization
func c() (i int, j string) {
	i = 10
	j = "rahul"
	return
}
