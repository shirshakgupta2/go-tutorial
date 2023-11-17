package main

import (
	"bufio"
	"fmt"
	"os"
	//"time"

	// "io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// This will hold information such
// as the Read and Write buffer size for our WebSocket connection:
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func homePage(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "Home Page")
}

func wsEndpoint(resp http.ResponseWriter, req *http.Request) {
	//This will determine whether or not an
	//incoming request from a different domain is allowed to connect,

	upgrader.CheckOrigin = func(resp *http.Request) bool { return true }

	/*will take in the Response Writer and the pointer to the HTTP Request
	and return us with a pointer to a WebSocket connection,*/

	// upgrade this connection to a WebSocket
	// connection
	websocket, err := upgrader.Upgrade(resp, req, nil)
	if err != nil {
		log.Println(err)
	}
	defer websocket.Close()
	// helpful log statement to show connections
	log.Println("client is succuesfully connected...")

	/*In order to send messages from our Go application to any connected WebSocket
	clients, we can utilize the ws.WriteMessage() function */

	//websocket is of type *websocket.com
	err = websocket.WriteMessage(1, []byte("Hi Client!"))

	if err != nil {
		log.Println(err)
		//return
	}

	/*call this reader() and writer() for now and it will take in a pointer to the WebSocket connection */
	go reader(websocket)
	writer(websocket)

	//time.Sleep(1*time.Hour)
	//fmt.Fprintf(resp, "we are hitting the end of the websocket connection")
}

/*function which will continually listen for any incoming messages sent through
that WebSocket connection. */

// define a reader which will listen for new messages being sent to our WebSocket endpoint
func writer(conn *websocket.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)

		input, err := reader.ReadBytes('\n')

		if err := conn.WriteMessage(1, input); err != nil {
			log.Println(err)
			return
		}

		if err != nil {
			log.Println(err)
		}
	}

}
func reader(conn *websocket.Conn) {

	for {

		// read in a message
		// ReadMessage methods to receive messages as a slice of bytes.
		_, p, err := conn.ReadMessage()

		if err != nil {
			log.Println(err)
			return
		}
		// print out that message
		fmt.Println(string(p))
		//time.Sleep(time.Duration(1) * time.Millisecond * 500)
		//err = conn.WriteMessage(1, []byte(string(input)))
		//writer(conn)

	}

}

// func reader(conn *websocket.Conn) {
// 	reader := bufio.NewReader(os.Stdin)
// 	for {
// 		// read in a message
// 		// ReadMessage methods to receive messages as a slice of bytes.
// 		messageType, p, err := conn.ReadMessage()
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}

// 		// print out that message
// 		fmt.Println(string(p))
// 		input, err := reader.ReadBytes('\n')
// 		// it is helper method for getting writer to write the msg
// 		if err := conn.WriteMessage(messageType, input); err != nil {
// 			log.Println(err)
// 			return
// 		}
// 		//err = conn.WriteMessage(1, []byte(string(input)))
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	}
// }

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	fmt.Println("Hello World")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":9090", nil))

}
