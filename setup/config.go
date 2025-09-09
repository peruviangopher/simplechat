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
)

type Config struct {
	rooms   int
	port    string
	secret  []byte
	userKey string
	mode    ServerMode
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

func LoadConfig() *Config {
	var roomsQuantity = flag.String("rooms", strconv.Itoa(defaultRooms), "Number of rooms available in app")
	var port = flag.String("port", defaultPort, "app http and socket port")
	var mode = flag.String("mode", string(defaultServerMode), fmt.Sprintf("server mode can be %s or %s, default %s", ServerModeAPI, ServerModeBot, defaultServerMode))

	flag.Parse() // parse the flags

	rooms, err := strconv.Atoi(*roomsQuantity)
	if err != nil {
		rooms = defaultRooms
	}

	return &Config{
		rooms:   rooms,
		port:    *port,
		secret:  secret,
		userKey: userKey,
		mode:    ServerMode(*mode),
	}
}
