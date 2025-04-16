package message

import (
	"context"
	"log"
	"time"

	"github.com/celtic93/chat-server/internal/model"

	sq "github.com/Masterminds/squirrel"
)

func (r *repo) Create(ctx context.Context, message *model.Message) error {
	log.Printf("repository.Message.Create started. Chat id: %d, User id: %d", message.ChatID, message.UserID)

	timeNow := time.Now()
	query, args, err := sq.Insert(messagesTable).
		Columns(chatIDColumn, userIDColumn, textColumn, createdAtColumn, updatedAtColumn).
		Values(message.ChatID, message.UserID, message.Text, timeNow, timeNow).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		log.Printf("failed to build query: %v", err)
		return err
	}

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		log.Printf("failed to insert message: %v", err)
		return err
	}

	log.Printf("repository.Message.Create ended. Chat id: %d, User id: %d", message.ChatID, message.UserID)

	return nil
}
