package main

import "fmt"

type Iphone13 struct {
	//attributes
	camera float32
	color  string
	ram    int
	rom    int
	price  float32
	size   float32
}

//method
func (i Iphone13) call(number string) {
	fmt.Println("calling..." + number)
}

func main() {
	//create instance of struct

	//empty instance

	a := Iphone13{}
	a.camera = 12.4
	a.color = "blue"
	a.call("9774688165")
	// non empty instance

	fmt.Println(a)
	b := Iphone13{}

}
