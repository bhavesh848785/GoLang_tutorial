package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator must not be zero")
	}
	return a/b, nil
}

/*
func divide(a, b float64) (float64, string) {
	if b == 0 {
		return 0, "denominator must not be zero"
	}
	return a/b, "nil"
}
*/

func main() {
	fmt.Println("satrting error handling in the fucntion")


	// error handling..........

	/*
	ans, err := divide(10,5)
	if err != nil {
		fmt.Println("error handing")
	}
	fmt.Println("division of two number is ", ans)
	*/

	//use underscore

	ans, _ := divide(10,0)
	fmt.Println("division of two number is", ans)

}