package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done() //signal that is goroutine are done
	fmt.Printf("worker %d started\n", i)
	fmt.Printf("worker %d end\n", i)
}

func main() {
	var wg sync.WaitGroup

	for i:=1; i<=3; i++{
		wg.Add(1) //increment the waitgroup counter
		go worker(i, &wg)
	}
	//wait for all worker
	wg.Wait()
	fmt.Println("worker task completed")
}