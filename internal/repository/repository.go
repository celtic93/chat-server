package repository

import (
	"context"

	"github.com/celtic93/chat-server/internal/model"
)

type ChatRepository interface {
	Create(ctx context.Context, chat *model.Chat) (int64, error)
	SendMessage(ctx context.Context, message *model.Message) error
	Delete(ctx context.Context, id int64) error
}
