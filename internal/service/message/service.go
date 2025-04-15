package message

import (
	"github.com/celtic93/chat-server/internal/repository"
	"github.com/celtic93/chat-server/internal/service"
)

type serv struct {
	messageRepository repository.MessageRepository
}

func NewService(messageRepository repository.MessageRepository) service.MessageService {
	return &serv{
		messageRepository: messageRepository,
	}
}
