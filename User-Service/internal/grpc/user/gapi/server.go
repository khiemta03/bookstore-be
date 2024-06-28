package gapi

import (
	"log"
	"net"

	db "github.com/khiemta03/bookstore-be/user-service/internal/database/sqlc"
	"github.com/khiemta03/bookstore-be/user-service/internal/grpc/user/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.UnimplementedUserServiceServer
	store *db.Store
}

// NewServer creates a new gRPC server
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}

	return server
}

func (server *Server) Run(address string) {
	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot create listener:", err)
	}
	defer listener.Close()

	log.Printf("user service server is running at: " + address)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start user service server:", err)
	}
}
