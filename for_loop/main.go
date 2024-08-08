package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("number is:", i)
	}

	counter:=0

	for{
		fmt.Println("infinite loop")
		counter++
		if counter==3{
			break;
		}
	}

	number := []int{11,22,33,44,55}
	for index, value := range number{
		fmt.Printf("index of number is %d, and value of number is %d\n", index, value)
	}

	data := "bhavesh v"
	for index, value:= range data{
		fmt.Printf("index of number is %d, and value of number is %c\n", index, value)
	}
}