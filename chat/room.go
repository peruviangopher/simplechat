package chat

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

type Room struct {

	// clients holds all current clients in this Room
	clients map[*client]bool

	// join is a channel for clients wishing to join the Room
	join chan *client

	// leave is a channel for clients wishing to leave the Room
	leave chan *client

	// forward is a channel that holds incoming messages that should be forwarded to the other clients
	forward chan []byte
}

func (r *Room) ServeHTTP(w http.ResponseWriter, req *http.Request, userName string) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}

	newClient := &client{
		name:    userName,
		socket:  socket,
		receive: make(chan []byte, messageBufferSize),
		room:    r,
	}

	r.join <- newClient
	defer func() { r.leave <- newClient }()

	go newClient.write()
	newClient.read()
}

// NewRoom
func NewRoom() *Room {
	return &Room{
		forward: make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
	}
}

func (r *Room) Run() {
	for {
		select {
		case newClient := <-r.join:
			r.clients[newClient] = true
		case removedClient := <-r.leave:
			delete(r.clients, removedClient)
			close(removedClient.receive)
		case msg := <-r.forward:
			for member := range r.clients {
				member.receive <- msg
			}
		}
	}
}
