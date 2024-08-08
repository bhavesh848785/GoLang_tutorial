package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("web request learning")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err!=nil{
		fmt.Println("not getting any response something wrong", err)
		return
	}

	defer res.Body.Close()
	fmt.Println("type of response:%T\n", res)

	//now read the response body
	data, error := ioutil.ReadAll(res.Body)
	if error!=nil{
		fmt.Println("error reading response", error)
		return
	}

	fmt.Println(string(data)) //without writing a string data display in byte form
}