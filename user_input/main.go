package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//read only before space like: bhavesh vaghela.....output is only vaghela
	
	/*fmt.Println("enter your name:")
	var name string;

	fmt.Scan(&name)
	fmt.Println("hey my name is:", name)
	*/


	//solution here....
	fmt.Println("enter your name:")

	reader := bufio.NewReader(os.Stdin)
	name2, _ := reader.ReadString('\n')
	fmt.Println("hello my name:", name2)

}