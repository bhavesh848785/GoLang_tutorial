package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4, 5}
	numbers = append(numbers, 6, 7, 8, 9, 10)
	fmt.Println(numbers)
	fmt.Printf("data types: %T\n", numbers)
	fmt.Println(len(numbers))

	name := []string{}
	fmt.Println(name)


	// here capacity is over then itself it making it double
	//use make function for lenght and capacity for arguments	
	data := make([]int, 4, 5)

	data = append(data, 4) // lenght =5 and capacity=5
	data = append(data, 5) // but here if length increase so capacity double itself, so lenght=6 then cap= 10

	fmt.Println(data)
	fmt.Println(len(data)) 
	fmt.Println(cap(data))
}