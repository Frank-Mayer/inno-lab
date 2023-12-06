package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"

	"github.com/go-vgo/robotgo"
	"github.com/gorilla/websocket"
)

type Message struct {
	Command string  `json:"command"`
	X       float64 `json:"x,omitempty"`
	Y       float64 `json:"y,omitempty"`
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
		err = json.Unmarshal(p, &message)
		if err != nil {
			fmt.Println(err)
		}

		switch message.Command {
		case "click":
			robotgo.MoveClick(int(int(math.Ceil(message.X))), int(int(math.Ceil(message.Y))))
			robotgo.TypeStrDelay("/imagine Professor lecturing", 500)
			err = robotgo.KeyTap("enter")
			if err != nil {
				fmt.Println(err)
			}
		default:
			fmt.Println("Unknown command " + message.Command)
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
