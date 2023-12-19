package logic

import (
	"github.com/Frank-Mayer/inno-lab/internal/schema"
	"github.com/charmbracelet/log"
	"github.com/go-vgo/robotgo"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

func Handle(conn *websocket.Conn, msg *schema.Message) {
	switch msg.Command {
	case schema.Command_RETURN_INPUT_POSITION:
		handleReturnInputPosition(conn, msg)
	case schema.Command_PROMPT_FINISHED:
		handlePromptFinished(msg)
	default:
		log.Error("Unhandeled Command", "Command", msg.Command, "Message", msg)
	}
}

func handleReturnInputPosition(conn *websocket.Conn, msg *schema.Message) {
	log.Debug("Handle", "Command", "RETURN_INPUT_POSITION", "X", msg.X, "Y", msg.Y)
	entry := queue.Q.Pop()
	if entry == nil {
		return
	}

	tempMsg := schema.Message{}
	tempMsg.Command = schema.Command_REGISTER_PROMPT
	tempMsg.Prompt = entry.Prompt
	b, err := proto.Marshal(&tempMsg)
	if err != nil {
		log.Error("Error marshaling protobuf message", "error", err)
		return
	}

	err = conn.WriteMessage(websocket.BinaryMessage, b)
	if err != nil {
		log.Error("Error writing message", "error", err)
		return
	}

	log.Debug("Sent", "Command", "REGISTER_PROMPT", "Prompt", msg.Prompt)

    robotgo.MoveClick(int(msg.X), int(msg.Y))
    robotgo.KeySleep = 100
    robotgo.TypeStr("/imagine ")
    robotgo.TypeStr(entry.Prompt)
    err = robotgo.KeyTap("enter")
    if err != nil {
        log.Error("Error typing", "key", "enter", "error", err)
        return
    }
}

func handlePromptFinished(msg *schema.Message) {
	log.Debug("Handle", "Command", "PROMPT_FINISHED", "Prompt", msg.Prompt, "Url", msg.Url)
}
