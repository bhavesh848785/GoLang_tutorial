package main

import "fmt"


func modifyValueByReference(num *int)  {
	*num = *num +20
}

func main() {
	//................................................
	num := 2
	ptr := &num

	fmt.Println("pointer contains:", ptr)
	fmt.Println("value show using pointer: ", *ptr)


	//..............................................
	var pointer *int// default pointer is nil
	if pointer == nil{
		fmt.Println("pointer is not assi	gned")
	}


	//.................................................
	value:= 10
	modifyValueByReference(&value)
	fmt.Println("value contains:", value)

}