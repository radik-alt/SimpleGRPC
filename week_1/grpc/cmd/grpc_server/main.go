package main

import (
	"context"
	"flag"
	"github.com/brianvoe/gofakeit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
	"week_1/grpc/internal/config"
	authGrpc "week_1/grpc/pkg/auth_v1"
	chatGrpc "week_1/grpc/pkg/chat_v1"
	desc "week_1/grpc/pkg/note_v1"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "prod.env", "path to config file")
}

type serverNote struct {
	desc.UnimplementedNoteV1Server
	//pool *pgxpool.Pool
}

type serverAuth struct {
	authGrpc.UnimplementedAuthV1Server
}

type serverChat struct {
	chatGrpc.UnimplementedChatV1Server
}

func (s *serverNote) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("Note id: %d", req.GetId())

	return &desc.GetResponse{
		Note: &desc.Note{
			Id: req.GetId(),
			Info: &desc.NoteInfo{
				Title:    gofakeit.BeerName(),
				Context:  gofakeit.IPv4Address(),
				Author:   gofakeit.Name(),
				IsPublic: gofakeit.Bool(),
			},
			CreatedAt: timestamppb.New(gofakeit.Date()),
			UpdatedAt: timestamppb.New(gofakeit.Date()),
		},
	}, nil
}

func (s *serverAuth) GetUser(ctx context.Context, req *authGrpc.GetUserRequest) (*authGrpc.User, error) {
	log.Printf("Auth id: %d", req.GetId())

	return &authGrpc.User{
		Id:        req.GetId(),
		Name:      gofakeit.BeerName(),
		Password:  gofakeit.Password(false, false, false, false, false, 0),
		CreatedAt: timestamppb.New(gofakeit.Date()),
	}, nil
}

func (s *serverChat) GetMessage(ctx context.Context, req *chatGrpc.GetMessageRequest) (*chatGrpc.ChatMessage, error) {
	log.Printf("Chat id: %d", req.GetId())

	return &chatGrpc.ChatMessage{
		Id:        req.GetId(),
		Message:   gofakeit.BeerName(),
		CreatedAt: timestamppb.New(gofakeit.Date()),
		UpdatedAt: timestamppb.New(gofakeit.Date()),
	}, nil
}

func main() {
	flag.Parse()
	//ctx := context.Background()

	// Считываем переменные окружения
	err := config.Load(configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	grpcConfig, err := config.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %v", err)
	}

	//pgConfig, err := config.NewPGConfig()
	//if err != nil {
	//	log.Fatalf("failed to get pg config: %v", err)
	//}

	lis, err := net.Listen("tcp", grpcConfig.GRPCAddress())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Создаем пул соединений с базой данных
	//pool, err := pgxpool.Connect(ctx, pgConfig.DSN())
	//if err != nil {
	//	log.Fatalf("failed to connect to database: %v", err)
	//}
	//defer pool.Close()

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterNoteV1Server(s, &serverNote{})
	authGrpc.RegisterAuthV1Server(s, &serverAuth{})
	chatGrpc.RegisterChatV1Server(s, &serverChat{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
