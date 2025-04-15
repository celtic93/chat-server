package chat

import (
	"context"
	"log"
)

func (s *serv) Delete(ctx context.Context, id int64) error {
	log.Printf("service.Chat.Delete started. Delete chat id: %d", id)

	err := s.chatRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	log.Printf("service.Chat.Delete ended. Deleted chat id: %d", id)

	return nil
}
