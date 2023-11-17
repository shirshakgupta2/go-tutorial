package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

// This will hold information such
// as the Read and Write buffer size for our WebSocket connection:
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type client struct {
	ID   int
	conn *websocket.Conn
}

func homePage(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "Home Page")
}

// CREATE VARIABLE WEBSOCKET

var clients []client
var i int = 0

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	//This will determine whether or not an
	//incoming request from a different domain is allowed to connect,

	upgrader.CheckOrigin = func(resp *http.Request) bool { return true }

	/*will take in the Response Writer and the pointer to the HTTP Request
	and return us with a pointer to a WebSocket connection,*/

	// upgrade this connection to a WebSocket
	// connection

	// INITIALIZE CONFIGs
	webConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	i = i + 1

	ranClient := client{
		ID:   i,
		conn: webConn,
	}
	clients = append(clients, ranClient)
	go reader(webConn)
	writer(webConn)

}
func writer(conn *websocket.Conn) {
	for {

		//_, msg, err := conn.ReadMessage()
		reader := bufio.NewReader(os.Stdin)

		input, err := reader.ReadBytes('\n')
		if err != nil {
			panic(err)
		}

		for _, client := range clients {
			if err = client.conn.WriteMessage(1, input); err != nil {
				return
			}
		}
	}
}

func reader(conn *websocket.Conn) {
	// LOOP IF CLIENT SEND TO SERVER
	for {
		// READ MESSAGE FROM BROWSER
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}

		// PRINT MESSAGE IN YOU CONSOLE TERMINAL
		fmt.Printf("send: %s\n", string(msg))

		// LOOP IF MESSAGE FOUND AND SEND AGAIN TO CLIENT FOR
		// WRITE IN YOU BROWSER

		for _, client := range clients {
			mssg := fmt.Sprintf("client ID %d messsage send %s", client.ID, msg)
			if err = client.conn.WriteMessage(1, []byte(mssg)); err != nil {
				return
			}
		}
		time.Sleep(1000 * time.Millisecond)

	}
}
func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	fmt.Println("Hello World")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":9090", nil))

}
