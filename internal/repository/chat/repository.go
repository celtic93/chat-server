package chat

import (
	"github.com/celtic93/chat-server/internal/repository"
	"github.com/jackc/pgx/v4/pgxpool"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repository.ChatRepository {
	return &repo{db: db}
}
