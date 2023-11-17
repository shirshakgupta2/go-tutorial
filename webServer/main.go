package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://lco.dev"

func main() {
	fmt.Println("LCO web request")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Respnse is of type %T", response)
	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content :=string(databytes)
	fmt.Println(content)
	

}
