package api

import (
	"log"
	"net"

	db "github.com/khiemta03/bookstore-be/book-service/internal/database/sqlc"
	pb "github.com/khiemta03/bookstore-be/book-service/internal/grpc/gen/book"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.UnimplementedBookServiceServer
	store *db.Store
}

// NewServer creates a new gRPC server
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}

	return server
}

func (server *Server) Run(address string) {
	grpcServer := grpc.NewServer()

	pb.RegisterBookServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot create listener:", err)
	}
	defer listener.Close()

	log.Printf("book service server is running at: " + address)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start book service server:", err)
	}
}