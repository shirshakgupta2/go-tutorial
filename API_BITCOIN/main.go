package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Define a struct to hold the JSON data returned from the API
type StockData struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Volume int     `json:"volume"`
	Name   string  `json:"name"`
}

func main() {
	// Define the stock symbol you want to retrieve data for
	symbol := "AAPL"

	// Make a GET request to the API endpoint
	response, err := http.Get("https://api.example.com/stock?symbol=" + symbol)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Unmarshal the JSON data into a struct
	var data StockData
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	// Print the data
	fmt.Println("Symbol:", data.Symbol)
	fmt.Println("Price:", data.Price)
	fmt.Println("Volume:", data.Volume)
	fmt.Println("Name:", data.Name)
}
