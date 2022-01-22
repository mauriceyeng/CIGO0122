package main

func main() {

}
func fact(number int) int {
	if number == 1 || number == 0 {
		return 1
	}
	result := number * fact(number-1)
	return result

}
