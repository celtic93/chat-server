package main

import (
	"fmt"
	"log"
	"net"

	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/celtic93/chat-server/internal/config"
	server "github.com/celtic93/chat-server/internal/grpc-server"
	desc "github.com/celtic93/chat-server/pkg/v1/chat"
)

func main() {
	cfg := config.MustLoad()

	conn, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
	if err != nil {
		log.Fatal(color.RedString("failed to serve grpc server: %v", err))
	}

	log.Print(color.GreenString("UserAPI grpc server listening on: %s", conn.Addr().String()))

	gsrv := grpc.NewServer()
	reflection.Register(gsrv)

	desc.RegisterChatV1Server(gsrv, &server.Server{})

	if err = gsrv.Serve(conn); err != nil {
		log.Fatal(color.RedString("failed to serve grpc server: %v", err))
	}
}
