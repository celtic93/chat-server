package message

import (
	"context"
	"log"

	"github.com/celtic93/chat-server/internal/model"
)

func (s *serv) Create(ctx context.Context, message *model.Message) error {
	log.Printf("service.Message.Create started. Chat id: %d, User id: %d", message.ChatID, message.UserID)

	err := s.messageRepository.Create(ctx, message)
	if err != nil {
		return err
	}

	log.Printf("service.Message.Create ended. Chat id: %d, User id: %d", message.ChatID, message.UserID)

	return nil
}
