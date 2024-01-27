package server

import (
	"net/http"
	"sync/atomic"
	"time"

	"github.com/Frank-Mayer/inno-lab/internal/keyboard"
	"github.com/Frank-Mayer/inno-lab/internal/schema"
	"github.com/charmbracelet/log"
	"github.com/go-vgo/robotgo"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	inputPosX = -1
	inputPosY = -1
)

var (
	promptCh     = make(chan string)
	urlCh        = make(chan string)
	expectsImage = atomic.Bool{}
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	// redirect to discord
	http.Redirect(w, r, "https://discord.com/channels/@me/1019912799977230416", http.StatusTemporaryRedirect)
}

func setInputPos(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error("Error upgrading websocket", "error", err)
		return
	}
	defer c.Close()

	for {
		_, b, err := c.ReadMessage()
		if err != nil {
			log.Error("Error reading message", "error", err)
			break
		}
		pos := schema.Pos{}
		err = proto.Unmarshal(b, &pos)
		if err != nil {
			log.Error("Error unmarshalling message", "error", err)
			break
		}
		inputPosX = int(pos.X)
		inputPosY = int(pos.Y)
	}
}

func sendImage(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error("Error upgrading websocket", "error", err)
		return
	}
	defer c.Close()

	for {
		_, b, err := c.ReadMessage()
		if err != nil {
			log.Error("Error reading message", "error", err)
			break
		}
		res := schema.Res{}
		err = proto.Unmarshal(b, &res)
		if err != nil {
			log.Error("Error unmarshalling message", "error", err)
			break
		}
		tm := time.Unix(res.Time, 0)
		log.Info("Received image",
			"time", tm.String(),
			"prompt", res.Prompt,
			"image", res.Src,
		)
		if expectsImage.Load() {
			urlCh <- res.Src
		}
	}
}

func processPrompt(prompt string) {
	inputPosX := inputPosX + 50
	inputPosY := inputPosY + 20
	log.Info("Processing prompt", "prompt", prompt)
	robotgo.KeySleep = 300
	robotgo.MouseSleep = 200
	log.Debug("move click", "x", inputPosX, "y", inputPosY)
	robotgo.Move(inputPosX, inputPosY)
	robotgo.MoveClick(inputPosX, inputPosY)
	robotgo.MoveClick(inputPosX, inputPosY)
	go func() {
		log.Debug("sleep")
		time.Sleep(1 * time.Second)
		log.Debug("type")
		keyboard.SendKeys("/imagine " + prompt)
		err := robotgo.KeyTap("enter")
		if err != nil {
			log.Error("Error sending enter", "error", err)
		}
	}()
}

func Init() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/set_input_pos", setInputPos)
	http.HandleFunc("/send_image", sendImage)

	// Start the server on port 8080
	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Error("Server error", "error", err)
			return
		}
	}()

	go func() {
		for {
			prompt := <-promptCh
			if inputPosX != -1 && inputPosY != -1 {
				log.Info("Input position set", "x", inputPosX, "y", inputPosY)
				processPrompt(prompt)
			} else {
				log.Warn("Input position not set")
			}
		}
	}()
}

func SentPrompt(prompt string) string {
	promptCh <- prompt
	expectsImage.Store(true)
	url := <-urlCh
	expectsImage.Store(false)
	return url
}
