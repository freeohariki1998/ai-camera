package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WebSocketHub struct {
	client    map[*websocket.Conn]bool
	broadcast chan []byte
}

func NewWebSocketHub() *WebSocketHub {
	return &WebSocketHub{
		client:    make(map[*websocket.Conn]bool),
		broadcast: make(chan []byte),
	}
}

func (hub *WebSocketHub) Run() {
	for {
		msg := <-hub.broadcast
		for client := range hub.client {
			client.WriteMessage(websocket.TextMessage, msg)
		}
	}
}
func (hub *WebSocketHub) HandleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	hub.client[conn] = true
}
