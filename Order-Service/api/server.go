package api

import (
	"fmt"
	"log"
	"net"

	"github.com/khiemta03/bookstore-be/order-service/internal/config"
	db "github.com/khiemta03/bookstore-be/order-service/internal/database/sqlc"
	"github.com/khiemta03/bookstore-be/order-service/internal/grpc/client/book"
	pb "github.com/khiemta03/bookstore-be/order-service/internal/grpc/gen/order"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.UnimplementedOrderServiceServer
	config            *config.Config
	store             db.Store
	bookServiceClient *book.BookServiceClient
}

// NewServer creates a new gRPC server
func NewServerWithConfig(config *config.Config) (*Server, error) {
	store, err := db.NewStore(config.DBDriver, config.DBSource)
	if err != nil {
		return nil, fmt.Errorf("cannot load config: %w", err)
	}

	server := &Server{
		store:             store,
		config:            config,
		bookServiceClient: book.NewBookServiceClient(config.BookServiceServerAddress),
	}

	return server, nil
}

func NewMockServer(store db.Store, bookServiceClient *book.BookServiceClient) *Server {
	server := &Server{
		store:             store,
		bookServiceClient: bookServiceClient,
	}

	return server
}

func (server *Server) Run() {
	grpcServer := grpc.NewServer()

	pb.RegisterOrderServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	address := fmt.Sprintf("%s:%s", server.config.ServerHost, server.config.ServerPort)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot create listener:", err)
	}
	defer listener.Close()

	log.Printf("order service server is running at: " + address)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start order service server:", err)
	}
}
