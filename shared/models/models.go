package models

import "time"

// User represents a user entity
type User struct {
	ID        string    `json:"id" db:"id"`
	Email     string    `json:"email" db:"email"`
	Name      string    `json:"name" db:"name"`
	Password  string    `json:"-" db:"password"`
	Avatar    string    `json:"avatar,omitempty" db:"avatar"`
	Role      string    `json:"role" db:"role"`
	Active    bool      `json:"active" db:"active"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Post represents a social post
type Post struct {
	ID        string    `json:"id" db:"id"`
	UserID    string    `json:"user_id" db:"user_id"`
	Content   string    `json:"content" db:"content"`
	MediaURL  string    `json:"media_url,omitempty" db:"media_url"`
	Likes     int       `json:"likes" db:"likes"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Message represents a chat message
type Message struct {
	ID             string    `json:"id" db:"id"`
	ConversationID string    `json:"conversation_id" db:"conversation_id"`
	SenderID       string    `json:"sender_id" db:"sender_id"`
	Content        string    `json:"content" db:"content"`
	Read           bool      `json:"read" db:"read"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}

// Conversation represents a chat conversation
type Conversation struct {
	ID        string    `json:"id" db:"id"`
	Name      string    `json:"name,omitempty" db:"name"`
	IsGroup   bool      `json:"is_group" db:"is_group"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
