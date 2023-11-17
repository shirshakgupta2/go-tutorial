package main

import "fmt"

func main() {
	//first we have to establish a db conn for that we have to call DataMigration() which is inside db-connection.go
	DataMigration()
	//then we have to call ROUTINGHANDLER()  which is inside into routing.go
	RountingHandler()
	fmt.Println("url: ", urlDSN)
}
