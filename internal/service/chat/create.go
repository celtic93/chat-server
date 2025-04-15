package chat

import (
	"context"
	"fmt"
	"log"

	"github.com/celtic93/chat-server/internal/model"
)

func (s *serv) Create(ctx context.Context, chat *model.Chat) (int64, error) {
	log.Printf("service.Chat.Create started")

	if len(chat.UserIDs) == 0 {
		log.Print("Error. Usernames are empty")
		return 0, fmt.Errorf("usernames can't be empty")
	}

	id, err := s.chatRepository.Create(ctx, chat)
	if err != nil {
		return 0, err
	}

	log.Printf("service.Chat.Create ended. Created chat id: %d", id)

	return id, nil
}
