package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/Frank-Mayer/inno-lab/internal/logic"
	"github.com/Frank-Mayer/inno-lab/internal/queue"
	"github.com/Frank-Mayer/inno-lab/internal/schema"
	"github.com/charmbracelet/log"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
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
		log.Error("Error upgrading websocket", "error", err)
		return
	}
	defer conn.Close()

	log.Info("Client connected", "remoteAddr", conn.RemoteAddr())

	wg := sync.WaitGroup{}

    // server to client
	wg.Add(1)
	go func() {
		defer wg.Done()
        for {
            if queue.Q.Lenght() == 0 {
                continue
            }
            entry := queue.Q.Peek()
            if entry == nil {
                continue
            }

            msg := schema.Message{}
            logic.GetInputPosition(&msg)
            b, err := proto.Marshal(&msg)
            if err != nil {
                log.Error("Error marshaling protobuf message", "error", err)
                return
            }

            err = conn.WriteMessage(websocket.BinaryMessage, b)
            if err != nil {
                log.Error("Error writing message", "error", err)
                return
            }
        }
	}()

    // client to server
	wg.Add(1)
	go func() {
		defer wg.Done()
        for {
            // Read message from the client
            _, b, err := conn.ReadMessage()
            if err != nil {
                log.Error("Error reading message", "error", err)
                return
            }

            msg := schema.Message{}
            err = proto.Unmarshal(b, &msg)
            if err != nil {
                log.Error("Error unmarshaling protobuf message", "error", err)
                return
            }
            logic.Handle(conn, &msg)
        }
	}()

	wg.Wait()
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	// redirect to discord
	http.Redirect(w, r, "https://discord.com/channels/@me/1019912799977230416", http.StatusTemporaryRedirect)
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)
	http.HandleFunc("/", handleIndex)

	// Start the server on port 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Error("Error starting server", "error", err)
		return
	}
	log.Info("Started server", "port", 8080)

    // wait for user input
    var input string
    for {
        _, err := fmt.Scanln(&input)
        if err != nil {
            log.Error("Error reading input", "error", err)
            return
        }
        queue.Q.Push(queue.QueueEntry{Prompt: input})
    }
}
