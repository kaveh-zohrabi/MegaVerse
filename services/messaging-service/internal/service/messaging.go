package service

import (
	"github.com/google/uuid"

	"megaverse.dev/services/messaging-service/internal/repository"
)

type MessagingService struct {
	repo *repository.MessagingRepository
}

type CreateConversationRequest struct {
	Name    string   `json:"name"`
	IsGroup bool     `json:"is_group"`
	Members []string `json:"members"`
}

type SendMessageRequest struct {
	Content string `json:"content"`
}

func NewMessagingService(repo *repository.MessagingRepository) *MessagingService {
	return &MessagingService{repo: repo}
}

func (s *MessagingService) CreateConversation(userID string, req *CreateConversationRequest) (*repository.Conversation, error) {
	id := uuid.New().String()

	if err := s.repo.CreateConversation(id, req.Name, req.IsGroup); err != nil {
		return nil, err
	}

	// Add creator as member
	if err := s.repo.AddMember(id, userID); err != nil {
		return nil, err
	}

	// Add other members
	for _, memberID := range req.Members {
		if memberID != userID {
			_ = s.repo.AddMember(id, memberID)
		}
	}

	return s.repo.GetConversation(id)
}

func (s *MessagingService) GetUserConversations(userID string) ([]repository.Conversation, error) {
	return s.repo.GetUserConversations(userID)
}

func (s *MessagingService) SendMessage(conversationID, senderID, content string) (*repository.Message, error) {
	isMember, err := s.repo.IsMember(conversationID, senderID)
	if err != nil || !isMember {
		return nil, err
	}

	msg := &repository.Message{
		ID:             uuid.New().String(),
		ConversationID: conversationID,
		SenderID:       senderID,
		Content:        content,
	}

	if err := s.repo.SaveMessage(msg); err != nil {
		return nil, err
	}

	return msg, nil
}

func (s *MessagingService) GetMessages(conversationID, userID string, limit, offset int) ([]repository.Message, error) {
	isMember, err := s.repo.IsMember(conversationID, userID)
	if err != nil || !isMember {
		return nil, err
	}

	// Mark messages as read
	_ = s.repo.MarkAsRead(conversationID, userID)

	return s.repo.GetMessages(conversationID, limit, offset)
}
