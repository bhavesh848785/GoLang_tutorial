package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("say hello world")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("after some time......")
}

func sayHi()  {
	fmt.Println("say hii world")
}

func main() {
	fmt.Println("learning goroutines")
	go sayHello() //independent run
	sayHi()
	time.Sleep(100 * time.Millisecond)
}