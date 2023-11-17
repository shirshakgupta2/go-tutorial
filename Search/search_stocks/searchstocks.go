package searchstock

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shirshak/search/constants"
)

func SearchStock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Database connected")
	tradingSymbol := r.URL.Query().Get("tradingsymbol")

	var stocklist []string

	query := "SELECT * FROM stocks_lists WHERE trading_symbol LIKE '%" + tradingSymbol + "%' or name like '%" + tradingSymbol + "%' order by (trading_symbol LIKE '" + tradingSymbol + "%' and name LIKE '" + tradingSymbol + "%') desc limit 20"

	err := constants.DataBase.Raw(query).Scan(&stocklist)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	if len(stocklist) != 0 {
		 
		json.NewEncoder(w).Encode(&stocklist)

	}else{
		fmt.Println("Request failed")
	}

}
