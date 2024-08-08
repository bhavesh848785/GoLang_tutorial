package main

import "fmt"

func main() {

	studentGrades := make(map[string]int)

	studentGrades["bhavesh"] = 100
	studentGrades["dev"] = 98
	studentGrades["harsh"] = 99

	fmt.Println(studentGrades)
	fmt.Println("marks of bhavesh:", studentGrades["bhavesh"])
	studentGrades["bhavesh"] = 97
	fmt.Println("marks of bhavesh:", studentGrades["bhavesh"])

	// delete(studentGrades, "bhavesh")
	// fmt.Println("marks of bhavesh:", studentGrades["bhavesh"])

	//check key exists or not 
	grades, exists := studentGrades["bhavesh"]
	fmt.Println("grade: ", grades)
	fmt.Println("exists:", exists)

	//for loop
	for index, value:= range studentGrades{
		fmt.Printf("key is %s, and marks is %d\n", index, value)
	}
}