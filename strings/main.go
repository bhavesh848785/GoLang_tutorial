package main

import (
	"fmt"
	"strings"
)

func main() {

	//split.......................................
	data := "apple,orage,banana,gvava"
	parts := strings.Split(data, ",")
	fmt.Println(parts)


	//count.......................................
	str := "one two three two two four"
	occurence := strings.Count(str, "two")
	fmt.Println("count:", occurence)

	//TrimSpace..................................

	str = "                   hello  world         "
	after := strings.TrimSpace(str)
	fmt.Println("after Trim:", after)


	//concate......................................

	str1 := "bhavesh"
	str2 := "vaghela"

	result := strings.Join([]string{str1, "kumar", str2}, " ")
	fmt.Println(result)
}