package main

import (
	"github.com/Frank-Mayer/inno-lab/internal/utils"
	"github.com/go-vgo/robotgo"
)

var (
	savePosX = utils.EnvInt("SAVE_POS_X")
	savePosY = utils.EnvInt("SAVE_POS_Y")
)

func main() {
	robotgo.Move(savePosX, savePosY)
	robotgo.Move(savePosX, savePosY)
	robotgo.Move(savePosX, savePosY)
}
