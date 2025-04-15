package chat

import (
	"github.com/celtic93/chat-server/internal/repository"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	chatsTable      string = "chats"
	chatsUsersTable string = "chats_users"

	IDColumn        string = "id"
	chatIDColumn    string = "chat_id"
	userIDColumn    string = "user_id"
	createdAtColumn string = "created_at"
	updatedAtColumn string = "updated_at"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repository.ChatRepository {
	return &repo{db: db}
}
