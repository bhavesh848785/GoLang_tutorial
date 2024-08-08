package main

import "fmt"

func main() {
	var name [5]string

	name[0] = "bhavesh"
	name[1] = "vaghela"
	fmt.Println(name)

	var numbers = [8]int{1,2,3,4,5}
	fmt.Println("number is:", numbers)
	fmt.Println(len(numbers))
}