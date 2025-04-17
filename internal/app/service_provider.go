package app

import (
	"context"
	"log"

	"github.com/celtic93/chat-server/internal/closer"
	"github.com/celtic93/chat-server/internal/config"
	"github.com/celtic93/chat-server/internal/repository"
	"github.com/celtic93/chat-server/internal/service"
	"github.com/jackc/pgx/v4/pgxpool"

	chatAPI "github.com/celtic93/chat-server/internal/api/chat"
	messageAPI "github.com/celtic93/chat-server/internal/api/message"
	chatRepository "github.com/celtic93/chat-server/internal/repository/chat"
	messageRepository "github.com/celtic93/chat-server/internal/repository/message"
	chatService "github.com/celtic93/chat-server/internal/service/chat"
	messageService "github.com/celtic93/chat-server/internal/service/message"
)

type serviceProvider struct {
	chatImplementation *chatAPI.Implementation
	chatRepository     repository.ChatRepository
	chatService        service.ChatService

	messageImplementation *messageAPI.Implementation
	messageRepository     repository.MessageRepository
	messageService        service.MessageService

	grpcConfig config.GRPCConfig

	pgConfig config.PGConfig
	db       *pgxpool.Pool
}

func (sp *serviceProvider) DBClient(ctx context.Context) *pgxpool.Pool {
	if sp.db == nil {
		pool, err := pgxpool.Connect(ctx, sp.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}

		err = pool.Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %v", err)
		}
		closer.Add(func() error {
			pool.Close()

			return nil
		})

		sp.db = pool
	}

	return sp.db
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (sp *serviceProvider) GRPCConfig() config.GRPCConfig {
	if sp.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %v", err)
		}

		sp.grpcConfig = cfg
	}

	return sp.grpcConfig
}

func (sp *serviceProvider) PGConfig() config.PGConfig {
	if sp.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %v", err)
		}

		sp.pgConfig = cfg
	}

	return sp.pgConfig
}

func (sp *serviceProvider) ChatImplementation(ctx context.Context) *chatAPI.Implementation {
	if sp.chatImplementation == nil {
		sp.chatImplementation = chatAPI.NewImplementation(sp.ChatService(ctx))
	}

	return sp.chatImplementation
}

func (sp *serviceProvider) ChatRepository(ctx context.Context) repository.ChatRepository {
	if sp.chatRepository == nil {
		sp.chatRepository = chatRepository.NewRepository(sp.DBClient(ctx))
	}

	return sp.chatRepository
}

func (sp *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if sp.chatService == nil {
		sp.chatService = chatService.NewService(sp.ChatRepository(ctx))
	}

	return sp.chatService
}

func (sp *serviceProvider) MessageImplementation(ctx context.Context) *messageAPI.Implementation {
	if sp.messageImplementation == nil {
		sp.messageImplementation = messageAPI.NewImplementation(sp.MessageService(ctx))
	}

	return sp.messageImplementation
}

func (sp *serviceProvider) MessageRepository(ctx context.Context) repository.MessageRepository {
	if sp.messageRepository == nil {
		sp.messageRepository = messageRepository.NewRepository(sp.DBClient(ctx))
	}

	return sp.messageRepository
}

func (sp *serviceProvider) MessageService(ctx context.Context) service.MessageService {
	if sp.messageService == nil {
		sp.messageService = messageService.NewService(sp.MessageRepository(ctx))
	}

	return sp.messageService
}
