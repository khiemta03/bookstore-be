package gapi

import (
	"fmt"
	"log"
	"net"
	"github.com/khiemta03/bookstore-be/authentication-service/internal/config"
	db "github.com/khiemta03/bookstore-be/authentication-service/internal/database/sqlc"
	"github.com/khiemta03/bookstore-be/authentication-service/internal/grpc/authentication/pb"
	user "github.com/khiemta03/bookstore-be/authentication-service/internal/grpc/user/gpai"
	errors "github.com/khiemta03/bookstore-be/authentication-service/pkg/error"
	"github.com/khiemta03/bookstore-be/authentication-service/pkg/token"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.UnimplementedAuthenticationServiceServer
	store             *db.Store
	config            *config.Config
	userServiceClient user.UserServiceClient
	tokenMaker        token.Maker
}

// NewServer creates a new gRPC server
func NewServer(store *db.Store, config *config.Config) (*Server, error) {
	jwtMaker, err := token.NewJWTMaker(config.TokenSecretKey)
	if err != nil {
		return nil, fmt.Errorf(errors.InternalServerError)
	}

	server := &Server{
		store:             store,
		config:            config,
		userServiceClient: user.NewUserServiceClient(config.UserServiceServerAddress),
		tokenMaker:        jwtMaker,
	}

	return server, nil
}

func (server *Server) Run(address string) {
	grpcServer := grpc.NewServer()

	pb.RegisterAuthenticationServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot create listener:", err)
	}
	defer listener.Close()

	log.Printf("authentication service server is running at: " + address)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start authentication service server:", err)
	}
}
