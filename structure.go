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
	old    Iphone12
}

//another structure Iphone12
type Iphone12 struct {
	camera float32
}

//method
func (i Iphone13) call(number string) {
	fmt.Println(i.camera)
	fmt.Println("calling..." + number)
}

func main() {
	//create instance of struct

	//empty instance: a

	a := Iphone13{}
	a.camera = 12.4
	a.color = "blue"
	a.call("9774688165")

	//print data of instance a
	fmt.Println(a)

	// non empty instance: b
	b := Iphone13{
		camera: 16,
		ram:    6,
		rom:    64,
		size:   1,
		color:  "white",
		old: Iphone12{
			camera: 20,
		},
	}
	fmt.Println(b)

}
