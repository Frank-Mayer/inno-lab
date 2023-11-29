package main

import (
	"fmt"
	"net/http"

	"github.com/go-vgo/robotgo"
	"github.com/gorilla/websocket"
)

type Message struct {
	Command string `json:"command"`
	X       int    `json:"x,omitempty"`
	Y       int    `json:"y,omitempty"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	fmt.Println("Client connected")

	for {
		// Read message from the client
		_, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		// Print the received message
		fmt.Printf("Received message: %s\n", p)

		// parse message
		var message Message
		err = conn.ReadJSON(&message)
		if err != nil {
			fmt.Println(err)
			return
		}

		switch message.Command {
		case "click":
			robotgo.Move(message.X, message.Y)
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)

	// Start the server on port 8080
	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
