package server

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	sq "github.com/Masterminds/squirrel"
	desc "github.com/celtic93/chat-server/pkg/v1/chat"
)

const (
	chatIDColumn    string = "chat_id"
	chatsUsersTable string = "chats_users"
	usernameColumn  string = "username"
)

type Server struct {
	desc.UnimplementedChatV1Server
	Pool *pgxpool.Pool
}

// Create: creates chat
func (s *Server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("server.Create Chat id: %s", req.GetUsernames())

	if len(req.GetUsernames()) == 0 {
		log.Print("usernames are empty")
		return nil, status.Error(codes.InvalidArgument, "usernames can't be empty")
	}

	var chatID int64
	tx, err := s.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		log.Printf("cannot start transaction: %v", err)
		return nil, status.Errorf(codes.Aborted, "cannot start transaction: %v", err)
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
			return
		}
		_ = tx.Commit(ctx)
	}()

	err = tx.QueryRow(ctx, "INSERT INTO chats DEFAULT VALUES RETURNING id;").Scan(&chatID)
	if err != nil {
		log.Printf("failed to insert chat: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to insert chat: %v", err)
	}

	for _, username := range req.GetUsernames() {
		createUserChatsQuery, args, err := sq.Insert(chatsUsersTable).
			Columns(chatIDColumn, usernameColumn).
			Values(chatID, username).
			PlaceholderFormat(sq.Dollar).
			ToSql()
		if err != nil {
			log.Printf("failed to build query: %v", err)
			return nil, status.Errorf(codes.Internal, "failed to build query: %v", err)
		}

		_, err = tx.Exec(ctx, createUserChatsQuery, args...)
		if err != nil {
			log.Printf("failed to insert chats_users: %v", err)
			return nil, status.Errorf(codes.Internal, "failed to insert chats_users: %v", err)
		}
	}

	return &desc.CreateResponse{
		Id: chatID,
	}, nil
}

// SendMessage: sends message to chat
func (s *Server) SendMessage(_ context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("server.SendMessage Chat id: %d", req.GetChatId())

	return &emptypb.Empty{}, nil
}

// Delete: deletes chat
func (s *Server) Delete(_ context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("server.Delete Chat id: %d", req.GetId())

	return &emptypb.Empty{}, nil
}
