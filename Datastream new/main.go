package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	//"fmt"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

// func main() {
// 	tcpConn, _ := net.Dial("tcp", c.path.Host+":443")
// 	cf := &tls.Config{Rand: rand.Reader}
// 	ssl := tls.Client(tcpConn, cf)

// 	c.clientConn = httputil.NewClientConn(ssl, nil)

// }

func crypto_detail(w http.ResponseWriter, r *http.Request) {
	//func crypto_detail() {

	// websocket client connection
	w.Header().Set("Content-Type", "application/json")
	c, _, err := websocket.DefaultDialer.Dial("wss://ws.coincap.io/trades/binance", nil)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	subscriptionMessage := map[string]interface{}{
		"event":   "subscribe",
		"channel": "trades",
		"symbol":  "binance",
	}
	message, err := json.Marshal(subscriptionMessage)
	if err != nil {
		log.Fatalf("marshal: %v", err)
	}
	err = c.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Fatalf("write: %v", err)
	}

	input := make(chan Trade) // 1️⃣
	var mutex = &sync.Mutex{}
	var trade []Trade
	go func() { // 2️⃣
		// read from the websocket
		for {
			_, message, err := c.ReadMessage() // 3️⃣
			if err != nil {
				break
			}
			// unmarshal the message from json to go
			var response Trade
			json.Unmarshal(message, &trade) // 4️⃣
			// send the trade to the channel which is in go format not in json format
			mutex.Lock()
			trade = append(trade, response)
			mutex.Unlock()
			input <- response // 5️⃣
		}
		close(input) // 6️⃣
	}()

	// cryptoCoin := make(chan Trade)

	// go func() {
	// 	for trade := range input {
	// 		// if trade.Base == "bitcoin" && trade.Quote == "tether" {
	// 		// 	bitcoin <- trade
	// 		// }
	// 		cryptoCoin <- trade
	// 	}
	// 	close(cryptoCoin)
	// }()

	// print the trades
	for trade := range input {
		//var objMap []map[string]interface{}
		detail, _ := json.MarshalIndent(trade, " ", "\t")
		// json.Unmarshal([]byte(detail),&objMap)
		//fmt.Println(string(detail))
		//w.Write([]byte(string(detail)))

		details := string(detail)
		json.NewEncoder(w).Encode(&details)
		time.Sleep(1 * time.Second)
	}

}

type Trade struct {
	Exchange  string  `json:"exchange"`
	Base      string  `json:"base"`
	Quote     string  `json:"quote"`
	Direction string  `json:"direction"`
	Price     float64 `json:"price"`
	Volume    int64   `json:"volume"`
	Timestamp int64   `json:"timestamp"`
	PriceUsd  float64 `json:"priceUsd"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/cryto", crypto_detail).Methods("GET")

	err := http.ListenAndServe(":9898", r)
	if err != nil {
		log.Fatal(err)
	}
	//crypto_detail()
}

// package main

// import (
// 	"encoding/json"
// 	"log"
// 	"net/http"
// 	"sync"

// 	"github.com/gorilla/websocket"
// )

// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// }

// type trade struct {
// 	Event   string                 `json:"event"`
// 	Channel string                 `json:"channel"`
// 	Symbol  string                 `json:"symbol"`
// 	Data    map[string]interface{} `json:"data"`
// }

// func main() {
// 	conn, _, err := websocket.DefaultDialer.Dial("wss://ws.coincap.io/trades/binance", nil)
// 	if err != nil {
// 		log.Fatalf("dial: %v", err)
// 	}
// 	defer conn.Close()

// 	http.HandleFunc("/trades", func(w http.ResponseWriter, r *http.Request) {
// 		var trades []trade
// 		var mutex = &sync.Mutex{}
// 		for {
// 			_, message, err := conn.ReadMessage()
// 			if err != nil {
// 				log.Fatalf("read: %v", err)
// 				break
// 			}

// 			var response trade
// 			err = json.Unmarshal(message, &response)
// 			if err != nil {
// 				log.Fatalf("unmarshal: %v", err)
// 				continue
// 			}

// 			mutex.Lock()
// 			trades = append(trades, response)
// 			mutex.Unlock()

// 			w.Header().Set("Content-Type", "application/json")
// 			json.NewEncoder(w).Encode(&trades)
// 		}
// 	})

// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }
