package chat

import (
	"fmt"

	"github.com/gorilla/websocket"
)

// Client represents a single chatting user.
type client struct {
	// name is the client user name
	name string

	// socket is the web socket for this client.
	socket *websocket.Conn

	// receive is a channel to receive messages from form other clients.
	receive chan []byte

	// room is the room this client is chatting in.
	room *Room
}

func (c *client) read() {
	defer c.socket.Close()

	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}

		formatMsg := fmt.Sprintf("%s: %s", c.name, msg)

		c.room.forward <- []byte(formatMsg)
	}
}

func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.receive {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}
