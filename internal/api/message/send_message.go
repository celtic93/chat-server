package message

import (
	"context"
	"log"

	"github.com/celtic93/chat-server/internal/api/message/converter"
	desc "github.com/celtic93/chat-server/pkg/v1/chat"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("api.Message.SendMessage started. Chat id: %d, User id: %d", req.ChatId, req.UserId)

	err := i.messageService.SendMessage(ctx, converter.ToMessageFromSendMessageRequest(req))
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Printf("api.Message.SendMessage ended. Chat id: %d, User id: %d", req.ChatId, req.UserId)

	return &emptypb.Empty{}, nil
}
