package main

import "fmt"

func main() {
	w := new(int)
	name := new(string)
	result, re := c(w, name)
	fmt.Println(result)
	fmt.Println(re)
	fmt.Println(*name)
	fmt.Println(*w)
}
func c(w *int, name *string) (i int, j string) {
	i = 10
	j = "maurice"
	*w = 100
	*name = "bot"
	return
}
