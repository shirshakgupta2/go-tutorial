package main

import (
	"fmt"
	"net/url"
)
const myurl = "https://lco.dev:3000/learn?coursename=reactjs&paymentrid=ghbj456ghb"
func main() {
	fmt.Println("Handling urls in golang")
	fmt.Println(myurl)

	//parsing the url
//parseing means that we are extracting some ins=fomation from the url
	result,_ := url.Parse(myurl)
	fmt.Println(result)
	
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Query)//gives the query adress
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	query_parameter :=result.Query()// stored in key value pairs in query string format  	
	fmt.Printf("the type of query parameters are %T\n",query_parameter) 	 	 	 
	fmt.Println(query_parameter["coursename"])


	for query,val := range query_parameter{
		fmt.Printf("the query parameter is: %v and the value is: %v \n",query , val)
	}


	// CONSTRUCTING A URL

	partsOfURL := &url.URL{
		Scheme: "http",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=hitesh",
	 }

	 anotherURL := partsOfURL.String()

	 fmt.Printf("another URL is: %v \n",anotherURL)
}