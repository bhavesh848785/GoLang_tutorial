package main

import (
	"fmt"
	"mylearning/variable"
)

func main() {
	fmt.Println("hello world")

	variable.Hello()

	var name string = "bhavesh"
	var s_name = "vaghela"
	fmt.Println(name)
	fmt.Println(s_name)

	var salary int = 50000
	var new_salaray = 70000
	fmt.Println(salary)
	fmt.Println(new_salaray)

	fmt.Println("full name: ", name, " ", s_name)

	var dimension float64 = 33.45
	fmt.Println(dimension)

	var decided bool = false
	fmt.Println(decided)

	const pi = 3.14
	fmt.Println(pi)


	//short cut:

	person := "bhavesh mera name"
	fmt.Println(person)
}
