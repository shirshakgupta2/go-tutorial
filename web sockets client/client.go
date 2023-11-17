package main

import (
	"bufio"
	"log"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

func receiveHandler(connection *websocket.Conn) {

	for {
		_, msg, err := connection.ReadMessage()
		if err != nil {
			log.Println("Error in receive:", err)
			return
		}
		log.Printf("Received: %s\n", msg)
	}
}

func main() {

	socketUrl := "ws://localhost:9090" + "/ws"
	conn, _, err := websocket.DefaultDialer.Dial(socketUrl, nil)
	if err != nil {
		log.Fatal("Error connecting to Websocket Server:", err)
	}
	defer conn.Close()
	go receiveHandler(conn)

	// Our main loop for the client
	// We send our relevant packets here
	reader := bufio.NewReader(os.Stdin)
	for {
		// taking the input fromthe users using bufio package and newReader func
		input, _ := reader.ReadBytes('\n')
		time.After(time.Duration(1) * time.Millisecond * 1000)
		// Send an echo packet every second
		err := conn.WriteMessage(websocket.TextMessage, []byte(string(input)))
		if err != nil {
			log.Println("Error during writing to websocket:", err)
			return
		}

	}
}
