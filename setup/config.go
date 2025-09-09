package setup

import (
	"flag"
	"fmt"
	"strconv"
)

var secret = []byte("secretsc")

type ServerMode string

const (
	ServerModeBot ServerMode = "bot"
	ServerModeAPI ServerMode = "api"
)

const (
	userKey           = "user"
	defaultRooms      = 2
	defaultPort       = ":8080"
	defaultServerMode = ServerModeAPI
	botAPIKey         = "5af9d12d7f8e20a6f6f48e9d6f93f58b"
	botAPIKeyName     = "apikey"
	defaultBotPort    = ":8081"
)

type Config struct {
	rooms         int
	port          string
	secret        []byte
	userKey       string
	mode          ServerMode
	botAPIKey     string
	botAPIKeyName string
	botPort       string
}

func (c *Config) Rooms() int {
	return c.rooms
}

func (c *Config) Port() string {
	return c.port
}

func (c *Config) Secret() []byte {
	return c.secret
}

func (c *Config) UserKey() string {
	return c.userKey
}

func (c *Config) ServerMode() ServerMode {
	return c.mode
}

func (c *Config) BotAPIKey() string {
	return c.botAPIKey
}

func (c *Config) BotAPIKeyName() string {
	return c.botAPIKeyName
}

func (c *Config) BotPort() string {
	return c.botPort
}

func LoadConfig() *Config {
	var roomsQuantity = flag.String("rooms", strconv.Itoa(defaultRooms), "Number of rooms available in app")
	var port = flag.String("port", defaultPort, "app http and socket port")
	var mode = flag.String("mode", string(defaultServerMode), fmt.Sprintf("server mode can be %s or %s, default %s", ServerModeAPI, ServerModeBot, defaultServerMode))
	var botPort = flag.String("botport", defaultBotPort, "app bot api port")

	flag.Parse() // parse the flags

	rooms, err := strconv.Atoi(*roomsQuantity)
	if err != nil {
		rooms = defaultRooms
	}

	return &Config{
		rooms:         rooms,
		port:          *port,
		secret:        secret,
		userKey:       userKey,
		mode:          ServerMode(*mode),
		botAPIKeyName: botAPIKeyName,
		botAPIKey:     botAPIKey,
		botPort:       *botPort,
	}
}
