package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
	_ "modernc.org/sqlite"
)

var db *sql.DB
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
	rooms      map[string]map[*Client]bool
	register   chan *Client
	unregister chan *Client
}

var hub = &Hub{
	clients:    make(map[*Client]bool),
	broadcast:  make(chan []byte, 256),
	rooms:      make(map[string]map[*Client]bool),
	register:   make(chan *Client),
	unregister: make(chan *Client),
}

type Message struct {
	ID             int64  `json:"id"`
	ConversationID string `json:"conversation_id"`
	SenderID       string `json:"sender_id"`
	SenderName     string `json:"sender_name,omitempty"`
	Content        string `json:"content"`
	CreatedAt      string `json:"created_at"`
}

type Conversation struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	IsGroup   bool   `json:"is_group"`
	UpdatedAt string `json:"updated_at"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Online   bool   `json:"online"`
}

func initDB() {
	var err error
	db, err = sql.Open("sqlite", "./messenger.db")
	if err != nil {
		log.Fatal(err)
	}

	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id TEXT PRIMARY KEY,
			username TEXT UNIQUE NOT NULL,
			password_hash TEXT NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS conversations (
			id TEXT PRIMARY KEY,
			name TEXT,
			is_group BOOLEAN DEFAULT FALSE,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS conversation_members (
			conversation_id TEXT,
			user_id TEXT,
			joined_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (conversation_id, user_id)
		)`,
		`CREATE TABLE IF NOT EXISTS messages (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			conversation_id TEXT,
			sender_id TEXT,
			content TEXT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)`,
	}

	for _, q := range queries {
		if _, err := db.Exec(q); err != nil {
			log.Fatal(err)
		}
	}
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

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		userID = fmt.Sprintf("user_%d", time.Now().UnixNano())
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{
		conn:   conn,
		userID: userID,
		send:   make(chan []byte, 256),
	}

	hub.register <- client

	go client.writePump()
	go client.readPump()

	notify := map[string]interface{}{
		"type":    "user_online",
		"user_id": userID,
	}
	data, _ := json.Marshal(notify)
	hub.broadcast <- data
}

func (c *Client) readPump() {
	defer func() {
		hub.unregister <- c
		c.conn.Close()

		notify := map[string]interface{}{
			"type":    "user_offline",
			"user_id": c.userID,
		}
		data, _ := json.Marshal(notify)
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
			createAndSendPrivateMessage(c.userID, msg.ReceiverID, msg.Content)
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
	var senderName string
	db.QueryRow("SELECT username FROM users WHERE id = ?", senderID).Scan(&senderName)
	if senderName == "" {
		senderName = senderID
	}

	result, err := db.Exec(
		"INSERT INTO messages (conversation_id, sender_id, content) VALUES (?, ?, ?)",
		convID, senderID, content,
	)
	if err != nil {
		log.Println(err)
		return
	}

	id, _ := result.LastInsertId()
	db.Exec("UPDATE conversations SET updated_at = CURRENT_TIMESTAMP WHERE id = ?", convID)

	msg := Message{
		ID:             id,
		ConversationID: convID,
		SenderID:       senderID,
		SenderName:     senderName,
		Content:        content,
		CreatedAt:      time.Now().Format(time.RFC3339),
	}

	data, _ := json.Marshal(map[string]interface{}{
		"type":    "message",
		"message": msg,
	})
	hub.broadcast <- data
}

func createAndSendPrivateMessage(senderID, receiverID, content) {
	var convID string
	err := db.QueryRow(`
		SELECT cm1.conversation_id FROM conversation_members cm1
		JOIN conversation_members cm2 ON cm1.conversation_id = cm2.conversation_id
		JOIN conversations c ON c.id = cm1.conversation_id
		WHERE cm1.user_id = ? AND cm2.user_id = ? AND c.is_group = FALSE
	`, senderID, receiverID).Scan(&convID)

	if err != nil {
		convID = fmt.Sprintf("dm_%s_%s", senderID, receiverID)
		db.Exec("INSERT OR IGNORE INTO conversations (id, is_group) VALUES (?, FALSE)", convID)
		db.Exec("INSERT OR IGNORE INTO conversation_members (conversation_id, user_id) VALUES (?, ?)", convID, senderID)
		db.Exec("INSERT OR IGNORE INTO conversation_members (conversation_id, user_id) VALUES (?, ?)", convID, receiverID)
	}

	saveAndBroadcast(senderID, convID, content)
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

	id := fmt.Sprintf("user_%d", time.Now().UnixNano())
	_, err := db.Exec("INSERT INTO users (id, username, password_hash) VALUES (?, ?, ?)",
		id, req.Username, req.Password)
	if err != nil {
		http.Error(w, `{"error":"username already exists"}`, 409)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"id": id, "username": req.Username})
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	var id, username string
	err := db.QueryRow("SELECT id, username FROM users WHERE username = ? AND password_hash = ?",
		req.Username, req.Password).Scan(&id, &username)
	if err != nil {
		http.Error(w, `{"error":"invalid credentials"}`, 401)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"id": id, "username": username})
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, username FROM users")
	if err != nil {
		http.Error(w, `{"error":"failed"}`, 500)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		rows.Scan(&u.ID, &u.Username)
		u.Online = true
		users = append(users, u)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"users": users})
}

func conversationsHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")

	rows, err := db.Query(`
		SELECT c.id, c.name, c.is_group, c.updated_at 
		FROM conversations c
		JOIN conversation_members cm ON c.id = cm.conversation_id
		WHERE cm.user_id = ?
		ORDER BY c.updated_at DESC
	`, userID)
	if err != nil {
		http.Error(w, `{"error":"failed"}`, 500)
		return
	}
	defer rows.Close()

	var convs []Conversation
	for rows.Next() {
		var c Conversation
		rows.Scan(&c.ID, &c.Name, &c.IsGroup, &c.UpdatedAt)
		convs = append(convs, c)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"conversations": convs})
}

func messagesHandler(w http.ResponseWriter, r *http.Request) {
	convID := r.URL.Query().Get("conversation_id")
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit == 0 {
		limit = 50
	}

	rows, err := db.Query(`
		SELECT m.id, m.conversation_id, m.sender_id, u.username, m.content, m.created_at
		FROM messages m
		LEFT JOIN users u ON m.sender_id = u.id
		WHERE m.conversation_id = ?
		ORDER BY m.created_at DESC
		LIMIT ?
	`, convID, limit)
	if err != nil {
		http.Error(w, `{"error":"failed"}`, 500)
		return
	}
	defer rows.Close()

	var msgs []Message
	for rows.Next() {
		var m Message
		rows.Scan(&m.ID, &m.ConversationID, &m.SenderID, &m.SenderName, &m.Content, &m.CreatedAt)
		msgs = append(msgs, m)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"messages": msgs})
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
}

func main() {
	initDB()
	go hub.run()

	http.HandleFunc("/ws", handleWebSocket)
	http.HandleFunc("/api/register", registerHandler)
	http.HandleFunc("/api/login", loginHandler)
	http.HandleFunc("/api/users", usersHandler)
	http.HandleFunc("/api/conversations", conversationsHandler)
	http.HandleFunc("/api/messages", messagesHandler)
	http.HandleFunc("/api/health", healthHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("🚀 MegaVerse Messenger running on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
