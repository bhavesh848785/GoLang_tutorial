package main

import "fmt"

func main() {
	
	
	num := 10

	if num > 5 && (num >6 || num >7) {
		fmt.Println("ya its correct")
	}else{
		fmt.Println("no its not correct")
	}
}