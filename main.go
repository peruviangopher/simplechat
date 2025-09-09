package main

import (
	"simplechat/bot"
	"simplechat/chatserver"
	"simplechat/setup"
)

func main() {
	cfg := setup.LoadConfig()

	switch cfg.ServerMode() {
	case setup.ServerModeAPI:
		chatserver.Run(cfg)
	case setup.ServerModeBot:
		bot.RunServer(cfg)
	}
}
