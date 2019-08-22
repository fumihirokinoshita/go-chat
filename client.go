package main

import (
	"github.com/gorilla/websocket"
)

// client represents a single user who is chatting.
type client struct {
	// socket is WebSocket fot this client
	socket *websocket.Conn
	// send is the channel through which messages are sent
	send chan []byte
	// room is the chatroom in which this client is participating.
	room *room
}
