package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	pb "github.com/khiemta03/bookstore-be/api-getway/internal/grpc/gen/authentication"
)

func (server *Server) validateTokenMiddleware(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(errors.New("authorization token is required")))
		ctx.Abort()
		return
	}

	_, err := server.client.ValidateToken(&pb.ValidateTokenRequest{
		Token: token,
	})

	if err != nil {
		ctx.JSON(err.Code, ErrorResponse(err.Error))
		ctx.Abort()
		return
	}

	ctx.Next()
}
