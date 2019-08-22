package main

type room struct {
	// forward is a channel that holds messages for forwarding to other clients.
	forward chan []byte
}
