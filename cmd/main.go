package main

import (
	"context"
	"log"
	"net"

	"github.com/fatih/color"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/celtic93/chat-server/internal/config"
	server "github.com/celtic93/chat-server/internal/grpc-server"
	desc "github.com/celtic93/chat-server/pkg/v1/chat"
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

	desc.RegisterChatV1Server(gsrv, &server.Server{Pool: pool})

	if err = gsrv.Serve(conn); err != nil {
		log.Fatal(color.RedString("failed to serve grpc server: %v", err))
	}
}
