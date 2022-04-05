package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// We'll need to define an Upgrader
// this will require a Read and Write buffer size
// This will hold information such as the Read and Write buffer size for our WebSocket connection
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

// define a reader which will listen for new messages being sent to our WebSocket endpoint
// We’ll call this reader() for now and it will take in a pointer to the WebSocket connection
// that we received from our call to upgrader.Upgrade:
func reader(conn *websocket.Conn) {
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()

		if err != nil {
			log.Println(err)
			return
		}

		// print out that message for clarity
		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello World")

	// This will determine whether or not an incoming request from a different domain is allowed to connect, and if it isn’t they’ll be hit with a CORS error.
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// upgrade this connection to a WebSocket connection
	// upgrader.Upgrade() function which will take in the Response Writer
	// and the pointer to the HTTP Request and return us with a pointer to a WebSocket connection,
	// or an error if it failed to upgrade.
	ws, err := upgrader.Upgrade(w, r, nil)

	// helpful log statement to show connections
	log.Println("Client Connected")

	err = ws.WriteMessage(1, []byte("Hi Client!"))

	if err != nil {
		log.Println(err)
	}

	// listen indefinitely for new messages coming
	// through on our WebSocket connection
	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	fmt.Println("Go WebSockets")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
