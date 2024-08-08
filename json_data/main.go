package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json: "name"`
	Age     int    `json: "age"`
	IsAdult bool   `json: "is_adult"`
}

func main() {

	person := Person{Name: "bhavesh", Age: 20, IsAdult: true}
	fmt.Println(person)

	//convert person into JSON Encoding (Marshalling)
	jsonData, err := json.Marshal(person)
	if err!=nil {
		fmt.Println("coversion error")
	}
	fmt.Println(string(jsonData))

	//decoding (unmarshalling)

	var personData Person
	error := json.Unmarshal(jsonData, &personData)
	if error!=nil {
		fmt.Println("error unmarshalling", error)
		return
	}
	fmt.Println(personData)
}