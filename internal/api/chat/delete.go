package chat

import (
	"context"
	"log"

	desc "github.com/celtic93/chat-server/pkg/v1/chat"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("api.Chat.Delete started. Delete chat: %d", req.GetId())
	err := i.chatService.Delete(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Printf("api.Chat.Delete ended. Deleted chat: %d", req.Id)

	return &emptypb.Empty{}, nil
}
