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

func (c *client) read() {
	for {
		if _, msg, err := c.socket.ReadMessage(); err == nil {
			c.room.forward <- msg
		} else {
			break
		}
	}
	c.socket.Close()
}

func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
	c.socket.Close()
}
