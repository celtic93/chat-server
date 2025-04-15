package message

import (
	"github.com/celtic93/chat-server/internal/repository"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	messagesTable string = "messages"

	chatIDColumn    string = "chat_id"
	userIDColumn    string = "user_id"
	textColumn      string = "text"
	createdAtColumn string = "created_at"
	updatedAtColumn string = "updated_at"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repository.MessageRepository {
	return &repo{db: db}
}
