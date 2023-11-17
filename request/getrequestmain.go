package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb video - LCO")
	PerformGetRequest()
	// PerformPostJsonRequest()
	// PerformPostFormRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:9090/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	//var responseString strings.Builder
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("err occured in reading response")
	}

	// byteCount, _ := responseString.Write(content)

	// fmt.Println("ByteCount is: ", byteCount)
	// fmt.Println(responseString.String())

	fmt.Println(content)
	fmt.Println(string(content))
}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:9090/post"

	//fake json payload

	requestBody := strings.NewReader(`
		{
			"coursename":"Let's go with golang",
			"price": 0,
			"platform":"learnCodeOnline.in"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:9090/postform"

	//formdata

	data := url.Values{}
	data.Add("firstname", "hitesh")
	data.Add("lastname", "choudhary")
	data.Add("email", "hitesh@go.dev")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}
