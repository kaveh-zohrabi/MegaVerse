package websocket

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	rooms      map[string]map[*Client]bool
	mu         sync.RWMutex
}

type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
	room string
	userID string
}

type Message struct {
	Type string      `json:"type"`
	Room string      `json:"room,omitempty"`
	Data interface{} `json:"data"`
}

func NewHub() *Hub {
	return &Hub{
		clients:   make(map[*Client]bool),
		broadcast: make(chan []byte, 256),
		rooms:     make(map[string]map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		msg := <-h.broadcast
		h.mu.RLock()
		for client := range h.clients {
			select {
			case client.send <- msg:
			default:
				close(client.send)
				delete(h.clients, client)
			}
		}
		h.mu.RUnlock()
	}
}

func (h *Hub) BroadcastToRoom(room string, msg []byte) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if clients, ok := h.rooms[room]; ok {
		for client := range clients {
			select {
			case client.send <- msg:
			default:
				close(client.send)
				delete(h.clients, client)
			}
		}
	}
}

func HandleWebSocket(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}

	userID := r.URL.Query().Get("user_id")
	room := r.URL.Query().Get("room")
	if room == "" {
		room = "global"
	}

	client := &Client{
		hub:    hub,
		conn:   conn,
		send:   make(chan []byte, 256),
		room:   room,
		userID: userID,
	}

	hub.mu.Lock()
	hub.clients[client] = true
	if hub.rooms[room] == nil {
		hub.rooms[room] = make(map[*Client]bool)
	}
	hub.rooms[room][client] = true
	hub.mu.Unlock()

	go client.writePump()
	go client.readPump()
}

func (c *Client) readPump() {
	defer func() {
		c.hub.mu.Lock()
		delete(c.hub.clients, c)
		if clients, ok := c.hub.rooms[c.room]; ok {
			delete(clients, c)
		}
		c.hub.mu.Unlock()
		c.conn.Close()
	}()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			break
		}

		msg := Message{
			Type: "message",
			Room: c.room,
			Data: map[string]interface{}{
				"user_id":  c.userID,
				"content":  string(message),
			},
		}

		msgBytes, _ := json.Marshal(msg)
		c.hub.BroadcastToRoom(c.room, msgBytes)
	}
}

func (c *Client) writePump() {
	defer c.conn.Close()

	for msg := range c.send {
		if err := c.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
}
