package main

import "fmt"

//interface- data type- abstract data type
//interface have only methods signature, although there can be empty interface also
/*
type interface_name interface{
	method_1
	method_2


}
*/

type SalaryCalculator interface {
	Calculate() int
}
type PermanentJob struct {
	basicPay int
}

func (p PermanentJob) Calculate() int {
	return p.basicPay
}

type ContractJob struct {
	basicPay int
}

// var salary SalaryCalculator
// salary can hold the data of multiple data types which comes with certain conditions
//salary should hold the values of types PermanentJob and ContractJob
//the condition is that types (i.e, PermanentJob and ContractJob),must implement all the methods of the interface
func main() {

	//method-1
	/*
		pj := PermanentJob{
			basicPay: 10,
		}
		var sc SalaryCalculator
		sc = pj
		total := sc.Calculate()
		fmt.Println(total)
	*/

	var sc SalaryCalculator
	sc = PermanentJob{
		basicPay: 10,
	}
	total := sc.Calculate()
	fmt.Println(total)
}
