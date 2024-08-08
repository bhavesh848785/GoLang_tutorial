package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct{
	UserID int `json:"userId"`
	Id int `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}

//perform get.......................................................................
func performGetRequest()  {
	fmt.Println("web request learning")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err!=nil{
		fmt.Println("not getting any response something wrong", err)
		return
	}
	defer res.Body.Close()

	//status code..............................................................
	if res.StatusCode != http.StatusOK {
		fmt.Println("error in getting respose:", res.Status)
		return
	}
	//now read the response body........................................................
	// data, error := ioutil.ReadAll(res.Body)
	// if error!=nil{
	// 	fmt.Println("error reading response", error)
	// 	return
	// }
	// fmt.Println(string(data)) //without writing a string data display in byte form

	var todo Todo
	erro := json.NewDecoder(res.Body).Decode(&todo)
	if erro!=nil {
		fmt.Println("error decoding:", erro)
		return
	}
	fmt.Println("TODO: ", todo)
}


//perform post..........................................................................
func performPostRequest()  {
	todo := Todo{
		UserID: 23,
		Title: "vaghela",
		Completed: true,
	}

	//convert todo struct to json
	json_data, err := json.Marshal(todo)
	if err!= nil {
		fmt.Println("error marshaling:", err)
		return
	} 

	//convert json to string
	json_string := string(json_data)

	//convert json to reader
	json_reader := strings.NewReader(json_string)
	my_url := "https://jsonplaceholder.typicode.com/todos"

	//send post request
	res, err := http.Post(my_url, "application/json", json_reader)

	if err!=nil {
		fmt.Println("error sending request:", err)
		return
	}
	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Response:", string(data))

	fmt.Println(res.StatusCode)

}

//perform put...........................................................................
func performPutRequest()  {
	todo := Todo{
		UserID: 23456,
		Title: "vaghela family",
		Completed: false,
	}

	//convert todo struct to json
	json_data, err := json.Marshal(todo)
	if err!= nil {
		fmt.Println("error marshaling:", err)
		return
	} 

	json_string := string(json_data)

	//convert json to reader
	json_reader := strings.NewReader(json_string)
	const my_url = "https://jsonplaceholder.typicode.com/todos/1"

	//put request create here
	req, err := http.NewRequest(http.MethodPut, my_url, json_reader)
	if err!=nil {
		fmt.Println("error while put method:", err)
		return
	}
	req.Header.Set("content-type", "application/json")

	//send the request
	client := http.Client{}
	res, erro := client.Do(req)
	if erro!=nil {
	     fmt.Println("error while sending request:", erro)
	}
	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("response:", string(data))
}


//perform delete............................................................

func performDeletRequest()  {
	const my_url = "https://jsonplaceholder.typicode.com/todos/1"

	//create delete request
	req, err := http.NewRequest(http.MethodDelete, my_url, nil)
	if err!=nil {
		fmt.Println("error while delete method:", err)
		return
	}

		//send the request
	client := http.Client{}
	res, erro := client.Do(req)
	if erro!=nil {
		fmt.Println("error while sending request:", erro)
	}
	defer res.Body.Close()
	
	fmt.Println("response:", res.Status)

}


//........................................................................
func main() {
	//performGetRequest()
	//performPostRequest()
	//performPutRequest()
	performDeletRequest()
}