package main

import (
	"context"
	"log"
	"net"

	"github.com/fatih/color"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	chatAPI "github.com/celtic93/chat-server/internal/api/chat"
	messageAPI "github.com/celtic93/chat-server/internal/api/message"
	"github.com/celtic93/chat-server/internal/config"
	chatRepository "github.com/celtic93/chat-server/internal/repository/chat"
 	chatService "github.com/celtic93/chat-server/internal/service/chat"
	messageRepository "github.com/celtic93/chat-server/internal/repository/message"
 	messageService "github.com/celtic93/chat-server/internal/service/message"
	chatDesc "github.com/celtic93/chat-server/pkg/v1/chat"
	messageDesc "github.com/celtic93/chat-server/pkg/v1/message"
)

func main() {
	err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	grpcConfig, err := config.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %v", err)
	}

	pgConfig, err := config.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to get pg config: %v", err)
	}

	ctx := context.Background()
	pool, err := pgxpool.Connect(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	conn, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatal(color.RedString("failed to serve grpc server: %v", err))
	}

	log.Print(color.GreenString("UserAPI grpc server listening on: %s", conn.Addr().String()))

	gsrv := grpc.NewServer()
	reflection.Register(gsrv)

	chatRepo := chatRepository.NewRepository(pool)
 	chatServ := chatService.NewService(chatRepo)
	chatDesc.RegisterChatV1Server(gsrv, chatAPI.NewImplementation(chatServ))

	msgRepo := messageRepository.NewRepository(pool)
 	msgServ := messageService.NewService(msgRepo)
	messageDesc.RegisterMessageV1Server(gsrv, messageAPI.NewImplementation(msgServ))

	if err = gsrv.Serve(conn); err != nil {
		log.Fatal(color.RedString("failed to serve grpc server: %v", err))
	}
}
