package main

import (
	"bufio"
	"fmt"
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
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}

		// PRINT MESSAGE IN YOU CONSOLE TERMINAL
		fmt.Printf("send: %s\n", string(msg))
		input, _ := reader.ReadBytes('\n')
		time.After(time.Duration(1) * time.Millisecond * 1000)

		// Send an echo packet every second

		if err := conn.WriteMessage(websocket.TextMessage, []byte(string(input)));err != nil {
			log.Println("Error during writing to websocket:", err)
			return
		}

	}
}
