package chat

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
)

func (r *repo) Delete(ctx context.Context, id int64) error {
	log.Printf("repository.Chat.Delete started. Delete chat id: %d", id)

	builderDelete := sq.Delete(chatsTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{IDColumn: id})

	query, args, err := builderDelete.ToSql()
	if err != nil {
		log.Print(err)
		return err
	}

	if _, err = r.db.Exec(ctx, query, args...); err != nil {
		log.Print(err)
		return err
	}

	log.Printf("repository.Chat.Delete ended. Deleted chat id: %d", id)

	return nil
}
