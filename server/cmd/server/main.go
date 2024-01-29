package main

import (
	"github.com/Frank-Mayer/inno-lab/internal/server"
	"github.com/Frank-Mayer/inno-lab/internal/ui"
	"github.com/Frank-Mayer/inno-lab/internal/utils"

	"github.com/charmbracelet/log"
)

func main() {

	log.SetLevel(log.DebugLevel)

	log.Info("Starting server...")
	server.Init()

	log.Info("Starting UI...")
	window := ui.Init()
	window.SetFullScreen(utils.EnvBool("FULLSCREEN"))
	window.ShowAndRun()
}
