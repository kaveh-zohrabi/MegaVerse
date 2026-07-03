package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var (
	users       = make(map[string]*User)
	messages    = make(map[string][]Message)
	convos      = make(map[string]*Conversation)
	members     = make(map[string][]string)
	mu          sync.RWMutex
	nextMsgID   int64 = 1
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	Online   bool   `json:"online"`
}

type Message struct {
	ID             int64  `json:"id"`
	ConversationID string `json:"conversation_id"`
	SenderID       string `json:"sender_id"`
	SenderName     string `json:"sender_name"`
	Content        string `json:"content"`
	CreatedAt      string `json:"created_at"`
}

type Conversation struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	IsGroup   bool   `json:"is_group"`
	UpdatedAt string `json:"updated_at"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type Client struct {
	conn   *websocket.Conn
	userID string
	send   chan []byte
}

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

var hub = &Hub{
	clients:    make(map[*Client]bool),
	broadcast:  make(chan []byte, 256),
	register:   make(chan *Client),
	unregister: make(chan *Client),
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

func getOrCreateConvID(userID1, userID2 string) string {
	if userID1 > userID2 {
		userID1, userID2 = userID2, userID1
	}
	convID := fmt.Sprintf("dm_%s_%s", userID1, userID2)

	mu.Lock()
	if _, exists := convos[convID]; !exists {
		convos[convID] = &Conversation{
			ID:        convID,
			IsGroup:   false,
			UpdatedAt: time.Now().Format(time.RFC3339),
		}
		members[convID] = []string{userID1, userID2}
	}
	mu.Unlock()

	return convID
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "user_id required", 400)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{conn: conn, userID: userID, send: make(chan []byte, 256)}
	hub.register <- client

	mu.Lock()
	if u, ok := users[userID]; ok {
		u.Online = true
	}
	mu.Unlock()

	go client.writePump()
	go client.readPump()

	data, _ := json.Marshal(map[string]interface{}{"type": "user_online", "user_id": userID})
	hub.broadcast <- data
}

func (c *Client) readPump() {
	defer func() {
		hub.unregister <- c
		c.conn.Close()
		data, _ := json.Marshal(map[string]interface{}{"type": "user_offline", "user_id": c.userID})
		hub.broadcast <- data
	}()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			break
		}

		var msg struct {
			Type           string `json:"type"`
			ConversationID string `json:"conversation_id"`
			Content        string `json:"content"`
			ReceiverID     string `json:"receiver_id"`
		}
		if err := json.Unmarshal(message, &msg); err != nil {
			continue
		}

		switch msg.Type {
		case "message":
			if msg.ConversationID != "" {
				saveAndBroadcast(c.userID, msg.ConversationID, msg.Content)
			}
		case "private_message":
			convID := getOrCreateConvID(c.userID, msg.ReceiverID)
			saveAndBroadcast(c.userID, convID, msg.Content)
		}
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

func saveAndBroadcast(senderID, convID, content string) {
	mu.Lock()
	senderName := "Unknown"
	if u, ok := users[senderID]; ok {
		senderName = u.Username
	}

	msg := Message{
		ID:             nextMsgID,
		ConversationID: convID,
		SenderID:       senderID,
		SenderName:     senderName,
		Content:        content,
		CreatedAt:      time.Now().Format(time.RFC3339),
	}
	nextMsgID++

	messages[convID] = append(messages[convID], msg)
	if c, ok := convos[convID]; ok {
		c.UpdatedAt = msg.CreatedAt
	}
	mu.Unlock()

	data, _ := json.Marshal(map[string]interface{}{"type": "message", "message": msg})
	hub.broadcast <- data
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	if req.Username == "" || req.Password == "" {
		http.Error(w, `{"error":"username and password required"}`, 400)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for _, u := range users {
		if u.Username == req.Username {
			http.Error(w, `{"error":"username already exists"}`, 409)
			return
		}
	}

	id := fmt.Sprintf("user_%d", time.Now().UnixNano())
	users[id] = &User{ID: id, Username: req.Username, Password: req.Password, Online: true}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": id, "username": req.Username})
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	mu.RLock()
	defer mu.RUnlock()

	for _, u := range users {
		if u.Username == req.Username && u.Password == req.Password {
			u.Online = true
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"id": u.ID, "username": u.Username})
			return
		}
	}

	http.Error(w, `{"error":"invalid credentials"}`, 401)
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	mu.RLock()
	defer mu.RUnlock()

	var userList []User
	for _, u := range users {
		userList = append(userList, *u)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"users": userList})
}

func messagesHandler(w http.ResponseWriter, r *http.Request) {
	convID := r.URL.Query().Get("conversation_id")

	mu.RLock()
	defer mu.RUnlock()

	msgs := messages[convID]
	if msgs == nil {
		msgs = []Message{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"messages": msgs})
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
}

func main() {
	go hub.run()

	http.HandleFunc("/ws", handleWebSocket)
	http.HandleFunc("/api/register", registerHandler)
	http.HandleFunc("/api/login", loginHandler)
	http.HandleFunc("/api/users", usersHandler)
	http.HandleFunc("/api/messages", messagesHandler)
	http.HandleFunc("/api/health", healthHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("🚀 MegaVerse Messenger running on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
