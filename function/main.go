package main

import "fmt"

func simpleFunction() {
	fmt.Println("simple function")
}

func add(a, b int) int {
	return a+b
}

func main()  {
	fmt.Println("function in golang")
	simpleFunction()
	ans:= add(3,4)
	fmt.Println("addition of two number:", ans)
}
