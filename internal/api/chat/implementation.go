package chat

import (
	"github.com/celtic93/chat-server/internal/service"
	desc "github.com/celtic93/chat-server/pkg/v1/chat"
)

type Implementation struct {
	desc.UnimplementedChatV1Server
	chatService service.ChatService
}

func NewImplementation(chatService service.ChatService) *Implementation {
	return &Implementation{
		chatService: chatService,
	}
}
