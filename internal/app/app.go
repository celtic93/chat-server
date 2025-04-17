package app

import (
	"context"
	"log"
	"net"

	"github.com/celtic93/chat-server/internal/closer"
	"github.com/celtic93/chat-server/internal/config"
	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	chatDesc "github.com/celtic93/chat-server/pkg/v1/chat"
	messageDesc "github.com/celtic93/chat-server/pkg/v1/message"
)

type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	return a.runGRPCServer()
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGRPCServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
func (a *App) initConfig(_ context.Context) error {
	err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()

	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	reflection.Register(a.grpcServer)
	chatDesc.RegisterChatV1Server(a.grpcServer, a.serviceProvider.ChatImplementation(ctx))
	messageDesc.RegisterMessageV1Server(a.grpcServer, a.serviceProvider.MessageImplementation(ctx))

	return nil
}

func (a *App) runGRPCServer() error {
	log.Print(color.GreenString("UserAPI grpc server listening on: %s", a.serviceProvider.GRPCConfig().Address()))

	conn, err := net.Listen("tcp", a.serviceProvider.GRPCConfig().Address())
	if err != nil {
		return err
	}

	err = a.grpcServer.Serve(conn)
	if err != nil {
		return err
	}

	return nil
}
