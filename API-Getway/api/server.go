package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/khiemta03/bookstore-be/api-getway/internal/config"
	"github.com/khiemta03/bookstore-be/api-getway/internal/grpc/client"
)

type Server struct {
	config *config.Config
	client *client.Client
	router *gin.Engine
}

// NewServer creates a new HTTP server
func NewServer(config *config.Config) *Server {
	router := gin.Default()

	server := &Server{
		client: client.NewClient(client.Config{
			UserServiceAddress:  config.UserServiceAddress,
			AuthServiceAddress:  config.AuthenticationServiceAddress,
			BookServiceAddress:  config.BookServiceAddress,
			OrderServiceAddress: config.OrderServiceAddress,
		}),
		router: router,
		config: config,
	}

	return server
}

func (server *Server) Start() error {
	address := fmt.Sprintf("%s:%s", server.config.ServerHost, server.config.ServerPort)
	return server.router.Run(address)
}

func ErrorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
