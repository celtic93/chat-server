package server

import (
	"context"
	"log"

	"github.com/brianvoe/gofakeit"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/celtic93/chat-server/pkg/v1/chat"
)

type Server struct {
	desc.UnimplementedChatV1Server
	Pool *pgxpool.Pool
}

// Create: creates chat
func (s *Server) Create(_ context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("server.Create Chat id: %s", req.GetUsernames())

	return &desc.CreateResponse{
		Id: gofakeit.Int64(),
	}, nil
}

// SendMessage: sends message to chat
func (s *Server) SendMessage(_ context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("server.Update Chat id: %s", req.GetFrom())

	return &emptypb.Empty{}, nil
}

// Delete: deletes chat
func (s *Server) Delete(_ context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("server.Delete Chat id: %d", req.GetId())

	return &emptypb.Empty{}, nil
}
