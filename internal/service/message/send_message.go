package message

import (
	"context"
	"log"

	"github.com/celtic93/chat-server/internal/model"
)

func (s *serv) SendMessage(ctx context.Context, message *model.Message) error {
	log.Printf("service.Message.SendMessage started. Chat id: %d, User id: %d", message.ChatID, message.UserID)

	err := s.messageRepository.SendMessage(ctx, message)
	if err != nil {
		return err
	}

	log.Printf("service.Message.SendMessage ended. Chat id: %d, User id: %d", message.ChatID, message.UserID)

	return nil
}
