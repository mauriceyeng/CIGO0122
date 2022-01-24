package main

import "fmt"

type person struct {
	Name string
	Age  int
}

func (p person) details() string {
	return fmt.Sprintf("%v(%v years)", p.Name, p.Age)
}
func main() {
	a := person{"Maurice", 21}
	fmt.Println("name:", a.Name, " age:", a.Age)
	fmt.Println(a.details())
}
