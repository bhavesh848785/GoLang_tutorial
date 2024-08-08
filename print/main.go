package main

import "fmt"

func main() {
	age := 25
	name := "bhavesh"
	height := 5.8234567

	fmt.Println("age:", age, "height:", height, "name:", name)

	fmt.Printf("age is %d\n", age)
	fmt.Printf("height is %.2f\n", height)
	fmt.Printf("type of age is %T\n", age)
	fmt.Printf("type of height is %T\n", height)

	fmt.Printf("name is %s", name)
}


