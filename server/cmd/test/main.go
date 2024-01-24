package main

import (
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	for {
		time.Sleep(1 * time.Second)
		go robotgo.Move(rand.Intn(1000), rand.Intn(1000))
	}
}
