package main

import (
	"fmt"
	"strconv"
)

func main() {

	//integer to float
	var num int = 12
	fmt.Println("number is", num)
	fmt.Printf("type of num is %T\n", num)

	var data float64 = float64(num)
	data = data + 12.12;
	fmt.Println("data:", data)
	fmt.Printf("types of data is %T\n", data)


	//integer to string
	number1 := 123
	str := strconv.Itoa(number1)
	fmt.Println("str is ", str)
	fmt.Printf("type %T\n", str)


	//string to integer
	number_string := "123"
	number_int, _ := strconv.Atoi(number_string)
	number_int = number_int + 556
	fmt.Println("integer is ", number_int)
	fmt.Printf("type %T\n", number_int)


	//string to float

	num_string := "3.14"
	num_float, _ := strconv.ParseFloat(num_string, 64)
	fmt.Println("float data: ", num_float)
	fmt.Printf("type: %T\n", num_float)

}