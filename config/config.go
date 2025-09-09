package config

import (
	"flag"
	"strconv"
)

var secret = []byte("secretsc")

const (
	userKey      = "user"
	defaultRooms = 2
	defaultPort  = ":8080"
)

type Config struct {
	rooms   int
	port    string
	secret  []byte
	userKey string
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

func LoadConfig() *Config {
	var roomsQuantity = flag.String("rooms", strconv.Itoa(defaultRooms), "Number of rooms available in app")
	var port = flag.String("port", defaultPort, "app http and socket port")
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
	}
}
