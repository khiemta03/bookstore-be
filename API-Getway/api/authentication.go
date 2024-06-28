package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	pb "github.com/khiemta03/bookstore-be/api-getway/internal/grpc/gen/authentication"
)

type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (server *Server) login(ctx *gin.Context) {
	var req loginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	user, err := server.client.Login(&pb.LoginRequest{
		Username:  req.Username,
		Password:  req.Password,
		UserAgent: ctx.Request.UserAgent(),
		ClientIp:  ctx.ClientIP(),
	})

	if err != nil {
		ctx.JSON(err.Code, ErrorResponse(err.Error))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

type registerRequest struct {
	Username         string `json:"username" binding:"required"`
	Password         string `json:"password" binding:"required"`
	RepeatedPassword string `json:"repeated_password" binding:"required"`
	Email            string `json:"email" binding:"required"`
	Address          string `json:"address" binding:"required"`
	FullName         string `json:"full_name" binding:"required"`
	Phone            string `json:"phone" binding:"required"`
}

func (server *Server) register(ctx *gin.Context) {
	var req registerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	user, err := server.client.Register(&pb.RegisterRequest{
		Username:         req.Username,
		Password:         req.Password,
		RepeatedPassword: req.RepeatedPassword,
		Email:            req.Email,
		Address:          req.Address,
		Fullname:         req.FullName,
		Phone:            req.Phone,
		UserAgent:        ctx.Request.UserAgent(),
		ClientIp:         ctx.ClientIP(),
	})

	if err != nil {
		ctx.JSON(err.Code, ErrorResponse(err.Error))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

type renewTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

func (server *Server) renewToken(ctx *gin.Context) {
	var req renewTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	token, err := server.client.RenewAccessToken(&pb.RenewAccessTokenRequest{
		RefreshToken: req.RefreshToken,
	})

	if err != nil {
		ctx.JSON(err.Code, ErrorResponse(err.Error))
		return
	}

	ctx.JSON(http.StatusOK, token)
}
