package main

import (
	"math/rand"
	"time"

	"github.com/Frank-Mayer/inno-lab/internal/server"
	"github.com/go-vgo/robotgo"
)

func main() {
	server.TypeStr("/imagine")
	for {
		time.Sleep(1 * time.Second)
		go robotgo.Move(rand.Intn(1000), rand.Intn(1000))
	}
}
