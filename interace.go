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

type Integer int

func (i Integer) print() {
	fmt.Println(i)
}
func (i Integer) Calculate() int {
	return int(i)
}

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

func (c ContractJob) Calculate() int {
	return c.basicPay
}

type FreelanceJob struct {
	basicPay int
}

func (f FreelanceJob) Calculate() int {
	return f.basicPay
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

	//method-2
	/*
		var sc SalaryCalculator
		sc = PermanentJob{
			basicPay: 10,
		}
		total := sc.Calculate()
		fmt.Println(total)
	*/

	//method-3

	j := Integer(1008)
	j.print()

	pj := PermanentJob{
		basicPay: 10,
	}
	cj := ContractJob{
		basicPay: 100,
	}
	fj := FreelanceJob{
		basicPay: 1000,
	}
	//var sc SalaryCalculator

	sc := []SalaryCalculator{pj, cj, fj} //an array of types
	totalIncome(sc)

}

func totalIncome(sc []SalaryCalculator) {
	total := 0
	for _, v := range sc {
		total = total + v.Calculate()
	}
	fmt.Println(total)
}
