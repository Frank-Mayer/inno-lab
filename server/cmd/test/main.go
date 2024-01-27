package main

import (
	"github.com/Frank-Mayer/inno-lab/internal/keyboard"
	"github.com/go-vgo/robotgo"
)

func main() {
	// type something
	keyboard.SendKeys("/Hello World!")

	// move mouse
	robotgo.MoveSmooth(100, 200, 1.0, 100.0)
}
