package message

import (
	"github.com/celtic93/chat-server/internal/service"
	desc "github.com/celtic93/chat-server/pkg/v1/chat"
)

type Implementation struct {
	desc.UnimplementedChatV1Server
	messageService service.MessageService
}

func NewImplementation(messageService service.MessageService) *Implementation {
	return &Implementation{
		messageService: messageService,
	}
}
