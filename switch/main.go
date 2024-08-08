package main

import "fmt"

func main() {
	day := 3

	switch day {
	case 1:
		fmt.Println("monday")
	case 2:
		fmt.Println("tuesday")
	case 3: 
	    fmt.Println("wednesday")
	default:
		fmt.Println("unknown day")
	}

	month := "jan"

	switch month {
	case "jan", "feb", "mar":
		fmt.Println("winter")
	case "apr", "may", "jun":
		fmt.Println("spring")
	default:
		fmt.Println("other season")
	}


	temp := 25

	switch{
	case temp < 0:
		fmt.Println("freezing")
	case temp >= 0 && temp < 10:
		fmt.Println("cold")
	case temp >= 10 && temp < 20:
		fmt.Println("cool")
	case temp >= 20 && temp < 30:
		fmt.Println("warm")
	default:
		fmt.Println("hot")
		
	}
}