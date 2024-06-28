package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	pb "github.com/khiemta03/bookstore-be/api-getway/internal/grpc/gen/user"
)

type getUserByIdRequest struct {
	Id string `uri:"id" binding:"required"`
}

func (server *Server) getUserById(ctx *gin.Context) {
	var req getUserByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	user, err := server.client.GetUserById(&pb.GetUserByIdRequest{
		Id: req.Id,
	})

	if err != nil {
		ctx.JSON(err.Code, ErrorResponse(err.Error))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

type listUsersRequest struct {
	Page    int32 `form:"page,min=1,default=1" binding:"required"`
	PerPage int32 `form:"per_page,min=1,default=10" binding:"required"`
}

func (server *Server) listUsers(ctx *gin.Context) {
	var req listUsersRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	userList, err := server.client.ListUsers(&pb.ListUsersRequest{
		Page:    req.Page,
		PerPage: req.PerPage,
	})

	if err != nil {
		ctx.JSON(err.Code, ErrorResponse(err.Error))
		return
	}

	ctx.JSON(http.StatusOK, userList)
}
