package main

import (
	"github.com/Frank-Mayer/inno-lab/internal/keyboard"
	"github.com/go-vgo/robotgo"
)

func main() {
	// type something
	// keyboard.SendKeys("/Hello World!")
	str := `https://firebasestorage.googleapis.com/v0/b/inno-lab-85f72.appspot.com/o/1706386572153939064.jpg?alt=media&token=747dc81e-6f2c-440f-b811-5502af2c094b`
	println(str)
	keyboard.SendKeys(str)

	// move mouse
	robotgo.MoveSmooth(100, 200, 1.0, 3.0)
}
