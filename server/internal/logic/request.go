package logic

import "github.com/Frank-Mayer/inno-lab/internal/schema"

func GetInputPosition(msg *schema.Message) {
    msg.Reset()
    msg.Command = schema.Command_GET_INPUT_POSITION
}

func RegisterPrompt(msg *schema.Message, prompt string) {
    msg.Reset()
    msg.Command = schema.Command_REGISTER_PROMPT
    msg.Prompt = prompt
}
