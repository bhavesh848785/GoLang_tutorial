package main

import (
	"fmt"
)


func add(a, b int) int{
	return a+b;
}

func main() {

	data:= add(5,5)
	fmt.Println(data)

	defer fmt.Println("first")
	defer fmt.Println("second")
	fmt.Println("third")
}

//when we add defer keyword those line end of the output
//if two line have a defer keyword that time maintain in stack (LIFO) 