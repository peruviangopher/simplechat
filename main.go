package main

import (
	"simplechat/bot"
	"simplechat/server"
	"simplechat/setup"
)

func main() {
	cfg := setup.LoadConfig()

	switch cfg.ServerMode() {
	case setup.ServerModeAPI:
		server.Run(cfg)
	case setup.ServerModeBot:
		bot.RunServer()
	}
}
