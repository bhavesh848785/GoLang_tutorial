package main

import (
	"fmt"
	"time"
)

func main() {
	current_time := time.Now()
	fmt.Println("current time:", current_time)
	fmt.Printf("type: %T\n", current_time)

	formated := current_time.Format("2006/01/02, 15:04:05")
	formated2 := current_time.Format("Mon, 2006/01/02, 03:04:05 PM")
	fmt.Println("now formated time: ", formated)
	fmt.Println("now formated time: ", formated2)


	//string to time formate:........................................
	layout_str := "2006-01-02"
	date_str := "2024-07-20"
	formated_date_time, _ := time.Parse(layout_str, date_str)
	fmt.Println(formated_date_time)

	//day increases

	new_date := current_time.Add(24 * time.Hour) // return new date and time
	fmt.Println("new updated time:", new_date)
	formated_new_date := new_date.Format("2006/01/02 Monday")
	fmt.Println("new updated formated time:", formated_new_date)

}