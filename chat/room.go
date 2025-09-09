package chat

import (
	"log"
	"net/http"
	"simplechat/setup"

	"github.com/gorilla/websocket"
)

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

type Room struct {
	roomID string

	// clients holds all current clients in this Room
	clients map[*client]bool

	// join is a channel for clients wishing to join the Room
	join chan *client

	// leave is a channel for clients wishing to leave the Room
	leave chan *client

	// forward is a channel that holds incoming messages that should be forwarded to the other clients
	forward chan []byte
}

func (r *Room) GetID() string {
	return r.roomID
}

func (r *Room) SendExternalMsg(msg []byte) {
	r.forward <- msg
}

func (r *Room) ServeHTTP(w http.ResponseWriter, req *http.Request, userName string, cfg *setup.Config) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}

	newClient := &client{
		cfg:     cfg,
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
func NewRoom(id string) *Room {
	return &Room{
		roomID:  id,
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
