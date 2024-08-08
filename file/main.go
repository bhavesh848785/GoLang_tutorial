package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

/*
// creating a new file..................................................

	file, err := os.Create("example.txt")
	if err!=nil {
		fmt.Println("error occur when we try to create a new file", err)
		return
	}
	defer file.Close()

	content := "first content of file: my name is bhavesh"
	// _ , error := io.WriteString(file, content+"\n")
	byte, error := io.WriteString(file, content+"\n")
	fmt.Println("byte written:", byte)

	if error!= nil {
		fmt.Println("error while writing a content into file:", error)
		return
	}

	fmt.Println("succesfully created file")

*/

//reading a example.txt file................................................

	file, err := os.Open("example.txt")
	if err!=nil {
		fmt.Println("error occur when we try to open a file", err)
		return
	}
	defer file.Close()

	//create buffere for reading a file content
	buffer := make([]byte, 1024)

	//read the file content into the buffer

	for{
		n, error := file.Read(buffer)
		if error == io.EOF {
			break
		}

		if error != nil{
			fmt.Println("error while reading the file", error)
			return
		}

		//process the read content

		fmt.Println(string(buffer[:n]))
	}


	// using ioutils......................................................
	con , er := ioutil.ReadFile("demo.txt")
	if er!=nil {
		fmt.Println("error occur:", er)
		return
	}
	fmt.Println(string(con))

}  