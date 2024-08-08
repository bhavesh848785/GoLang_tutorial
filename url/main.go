package main

import (
	"fmt"
	"net/url"
)

func main() {
	myURL := "https://example.com:8080/path/to/resource?key1=value1&key2=value2"
	fmt.Printf("type: %T", myURL)

	parsedURL, err := url.Parse(myURL)
	if err != nil{
		fmt.Println("can't parse url", err)
	} 
	fmt.Printf("type: %T", parsedURL)

	fmt.Println("scheme:", parsedURL.Scheme)
	fmt.Println("Host:", parsedURL.Host)
	fmt.Println("path:", parsedURL.Path)
	fmt.Println("rowQuery:", parsedURL.RawQuery)

	//modifying URL components
	parsedURL.Path = "/newPath"
	parsedURL.RawQuery = "username=bhavesh"
	
	newURL := parsedURL.String()
	fmt.Println("new url:", newURL)
}