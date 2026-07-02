package repository

import (
	"database/sql"
	"time"
)

type Conversation struct {
	ID        string
	Name      string
	IsGroup   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Message struct {
	ID             string
	ConversationID string
	SenderID       string
	Content        string
	Read           bool
	CreatedAt      time.Time
}

type MessagingRepository struct {
	db *sql.DB
}

func NewMessagingRepository(db *sql.DB) *MessagingRepository {
	return &MessagingRepository{db: db}
}

func (r *MessagingRepository) CreateConversation(id, name string, isGroup bool) error {
	_, err := r.db.Exec(
		"INSERT INTO conversations (id, name, is_group) VALUES (?, ?, ?)",
		id, name, isGroup,
	)
	return err
}

func (r *MessagingRepository) AddMember(conversationID, userID string) error {
	_, err := r.db.Exec(
		"INSERT IGNORE INTO conversation_members (conversation_id, user_id) VALUES (?, ?)",
		conversationID, userID,
	)
	return err
}

func (r *MessagingRepository) GetUserConversations(userID string) ([]Conversation, error) {
	rows, err := r.db.Query(
		`SELECT c.id, COALESCE(c.name,''), c.is_group, c.created_at, c.updated_at 
		 FROM conversations c 
		 JOIN conversation_members cm ON c.id = cm.conversation_id 
		 WHERE cm.user_id = ? 
		 ORDER BY c.updated_at DESC`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var conversations []Conversation
	for rows.Next() {
		var conv Conversation
		if err := rows.Scan(&conv.ID, &conv.Name, &conv.IsGroup, &conv.CreatedAt, &conv.UpdatedAt); err != nil {
			return nil, err
		}
		conversations = append(conversations, conv)
	}
	return conversations, nil
}

func (r *MessagingRepository) GetConversation(id string) (*Conversation, error) {
	conv := &Conversation{}
	err := r.db.QueryRow(
		"SELECT id, COALESCE(name,''), is_group, created_at, updated_at FROM conversations WHERE id = ?",
		id,
	).Scan(&conv.ID, &conv.Name, &conv.IsGroup, &conv.CreatedAt, &conv.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return conv, nil
}

func (r *MessagingRepository) IsMember(conversationID, userID string) (bool, error) {
	var count int
	err := r.db.QueryRow(
		"SELECT COUNT(*) FROM conversation_members WHERE conversation_id = ? AND user_id = ?",
		conversationID, userID,
	).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *MessagingRepository) SaveMessage(msg *Message) error {
	_, err := r.db.Exec(
		"INSERT INTO messages (id, conversation_id, sender_id, content) VALUES (?, ?, ?, ?)",
		msg.ID, msg.ConversationID, msg.SenderID, msg.Content,
	)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(
		"UPDATE conversations SET updated_at = NOW() WHERE id = ?",
		msg.ConversationID,
	)
	return err
}

func (r *MessagingRepository) GetMessages(conversationID string, limit, offset int) ([]Message, error) {
	rows, err := r.db.Query(
		`SELECT id, conversation_id, sender_id, content, read_status, created_at 
		 FROM messages WHERE conversation_id = ? 
		 ORDER BY created_at DESC LIMIT ? OFFSET ?`,
		conversationID, limit, offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var msg Message
		if err := rows.Scan(&msg.ID, &msg.ConversationID, &msg.SenderID, &msg.Content, &msg.Read, &msg.CreatedAt); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}
	return messages, nil
}

func (r *MessagingRepository) MarkAsRead(conversationID, userID string) error {
	_, err := r.db.Exec(
		`UPDATE messages SET read_status = TRUE 
		 WHERE conversation_id = ? AND sender_id != ? AND read_status = FALSE`,
		conversationID, userID,
	)
	return err
}
