package chat

import (
	"context"
	"log"

	"github.com/celtic93/chat-server/internal/api/chat/converter"
	desc "github.com/celtic93/chat-server/pkg/v1/chat"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("api.Chat.Create started. Create chat")
	id, err := i.chatService.Create(ctx, converter.ToChatFromCreateRequest(req))
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Printf("api.Chat.Create ended. Created chat with id: %d", id)

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
