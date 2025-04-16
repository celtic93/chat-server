package chat

import (
	"context"
	"log"
	"time"

	"github.com/celtic93/chat-server/internal/model"
	"github.com/jackc/pgx/v4"

	sq "github.com/Masterminds/squirrel"
)

func (r *repo) Create(ctx context.Context, chat *model.Chat) (int64, error) {
	log.Printf("repository.Chat.Create started")

	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		log.Printf("cannot start transaction: %v", err)
		return 0, err
	}

	var chatID int64
	defer func() {
		if err != nil {
			log.Printf("failed to insert chat and chats_users: %v", err)
			_ = tx.Rollback(ctx)
			return
		}
		log.Print("insert chat and chats_users")
		_ = tx.Commit(ctx)

		log.Printf("repository.Chat.Create ended. Created chat id: %d", chatID)
	}()

	timeNow := time.Now()
	builderInsert := sq.Insert(chatsTable).
		PlaceholderFormat(sq.Dollar).
		Columns(createdAtColumn, updatedAtColumn).
		Values(timeNow, timeNow).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Print(err)
		return 0, err
	}

	err = tx.QueryRow(ctx, query, args...).Scan(&chatID)
	if err != nil {
		log.Printf("failed to insert chat: %v", err)
		return 0, err
	}

	for _, userID := range chat.UserIDs {
		createUserChatsQuery, args, err := sq.Insert(chatsUsersTable).
			Columns(chatIDColumn, userIDColumn, createdAtColumn, updatedAtColumn).
			Values(chatID, userID, timeNow, timeNow).
			PlaceholderFormat(sq.Dollar).
			ToSql()
		if err != nil {
			log.Printf("failed to build query: %v", err)
			return 0, err
		}

		_, err = tx.Exec(ctx, createUserChatsQuery, args...)
		if err != nil {
			log.Printf("failed to insert chats_users: %v", err)
			return 0, err
		}
	}

	return chatID, nil
}
