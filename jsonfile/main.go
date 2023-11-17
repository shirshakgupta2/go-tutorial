package main

import (
	"encoding/json"
	"fmt"
)


// type course struct {
// 	Name     string
// 	Price    int
// 	Platfrom string
// 	Password string
// 	Tags     []string
// }
type course struct {
	Name     string `json:"coursename"`// these are just like aliases for the struct define
	Price    int  `json:"price"`
	Website string 	`json:"platfrom"`
	Password string  `json:"-"`// to secure the data where the json should not show the data
	Tags     []string `json:"tags,omitempty"`//if the feild is empty then do not through the feild 
}
func main() {
	fmt.Println("welcomme to json handling...")

	//EncodeJson()

	DecodeJson()
}


// CONVERTING EVERY DATA OF COURSE TO JSON
func EncodeJson() {
	lcoCourses := []course{ //this is storing in the slice of courses
		// hard coding the values in slice of course which is storing course kind of data
		{"ReactJs Bootcamp", 299, "Lco.in", "abc123", []string{"ReactJs", "web-dev"}},
		{"MERN Bootcamp", 199, "Lco.in", "bcd123", []string{"flwd", "js"}},
		{"Anugukr Bootcamp", 99, "Lco.in", "shi123", nil},
	}

	//package this data as JSON data

	//finalJson, err := json.Marshal(lcoCourses)//intendation in this is not good
	finalJson, err := json.MarshalIndent(lcoCourses,"","\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)

}


func DecodeJson()  {
	// whenever data is coming web ot will be in a byte format of slice then we convert in string or json
	jsonDataFromWeb :=[]byte(`// we hard coded the json data which will be coming from web
	{
			"coursename": "ReactJs Bootcamp",
			"price": 299,
			"platfrom": "Lco.in",
			"tags": ["ReactJs","web-dev"]
	}
	`)

	var lcoCourse course


	checkvalid:=json.Valid(jsonDataFromWeb)// json  package has a valid function which validates the json data

	if checkvalid{
		fmt.Println("JSON data was valid")
		json.Unmarshal(jsonDataFromWeb,&lcoCourse)// UNMARSAL CONTAIN A DATA AND INTERFACE AND WE ARE PASSING ON THE REFRENCE 
		// BECOZ IF IN METHOD WE PASS THE VALUE THEN THE COPY OF THE VALUE WILL BE TRANSFERED 

		fmt.Printf("%#v\n",lcoCourse)

	}else{
		fmt.Println("JSON WAS NOT VALID")
	}

fmt.Println("")
fmt.Println("")
	// some case where you just want to add data to key value pairs
	

	var myOnlineData map[string]interface{}//when we create a map for online json the first value 
	//is key value pair that is sure but the value that comes can be  int,string,array anything

	json.Unmarshal(jsonDataFromWeb,&myOnlineData)

	fmt.Printf("%#v\n\n",myOnlineData)


	for key,value := range myOnlineData {
		fmt.Printf("key is %v and value is %v and TYPE is: %T \n",key,value,value)
	}

}
