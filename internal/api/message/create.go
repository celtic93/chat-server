package message

import (
	"context"
	"log"

	"github.com/celtic93/chat-server/internal/api/message/converter"
	desc "github.com/celtic93/chat-server/pkg/v1/message"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*emptypb.Empty, error) {
	log.Printf("api.Message.Create started. Chat id: %d, User id: %d", req.ChatId, req.UserId)

	err := i.messageService.Create(ctx, converter.ToMessageFromCreateRequest(req))
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Printf("api.Message.Create ended. Chat id: %d, User id: %d", req.ChatId, req.UserId)

	return &emptypb.Empty{}, nil
}
