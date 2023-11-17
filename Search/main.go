package main

import (
	"github.com/shirshak/search/constants"
	"github.com/shirshak/search/router"

)

func main() {
	constants.InitDataBase()
	router.InitRouter()
	


}
