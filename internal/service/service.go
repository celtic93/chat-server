package service

import (
	"context"

	"github.com/celtic93/chat-server/internal/model"
)

type ChatService interface {
	Create(ctx context.Context, chat *model.Chat) (int64, error)
	Delete(ctx context.Context, id int64) error
}

type MessageService interface {
	Create(ctx context.Context, message *model.Message) error
}
