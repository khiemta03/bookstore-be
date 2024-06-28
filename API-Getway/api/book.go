package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	pb "github.com/khiemta03/bookstore-be/api-getway/internal/grpc/gen/book"
	"github.com/khiemta03/bookstore-be/api-getway/internal/util"
)

type getBookRequest struct {
	Id string `uri:"id" binding:"required"`
}

func (server *Server) getBook(ctx *gin.Context) {
	var req getBookRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	book, err := server.client.GetBook(&pb.GetBookRequest{
		Id: req.Id,
	})

	if err != nil {
		ctx.JSON(err.Code, ErrorResponse(err.Error))
		return
	}

	ctx.JSON(http.StatusOK, book)
}

type listBooksRequest struct {
	Page    int32 `form:"page,min=1,default=1" binding:"required"`
	PerPage int32 `form:"per_page,min=1,default=10" binding:"required"`
}

func (server *Server) listBooks(ctx *gin.Context) {
	var req listBooksRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	bookList, err := server.client.ListBooks(&pb.ListBooksRequest{
		Page:    req.Page,
		PerPage: req.PerPage,
	})

	if err != nil {
		ctx.JSON(err.Code, ErrorResponse(err.Error))
		return
	}

	ctx.JSON(http.StatusOK, bookList)
}

type addBookRequest struct {
	Title           string  `json:"title" binding:"required"`
	FullTitle       string  `json:"full_title" binding:"required"`
	Pubisher        string  `json:"publisher" binding:"required"`
	PublicationDate string  `json:"publication_date"`
	Isbn            string  `json:"isbn" binding:"required"`
	Description     string  `json:"description" binding:"required"`
	Price           float64 `json:"price" binding:"required"`
	StockQuantity   int32   `json:"stock_quantity" binding:"required"`
	FrontCoverImage string  `json:"front_cover_image" binding:"required"`
	BackCoverImage  string  `json:"back_cover_image" binding:"required"`
	Authors         string  `json:"authors" binding:"required"`
}

func (server *Server) addBook(ctx *gin.Context) {
	var req addBookRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	convertedDate, err := util.ConvertStringToDate(req.PublicationDate)
	if err != nil {
		ctx.JSON(err.Code, ErrorResponse(err.Error))
		return
	}

	book, err := server.client.AddNewBook(&pb.AddNewBookRequest{
		Title:           req.Title,
		FullTitle:       req.FullTitle,
		Publisher:       req.Pubisher,
		PublicationDate: convertedDate,
		Isbn:            req.Isbn,
		Description:     req.Description,
		Price:           req.Price,
		StockQuantity:   req.StockQuantity,
		FrontCoverImage: req.FrontCoverImage,
		BackCoverImage:  req.BackCoverImage,
		Authors:         util.ParseStringToList(req.Authors, ","),
	})

	if err != nil {
		ctx.JSON(err.Code, ErrorResponse(err.Error))
		return
	}

	ctx.JSON(http.StatusOK, book)
}

type updateBookRequest struct {
	Id              string  `json:"id" binding:"required"`
	Title           string  `json:"title"`
	FullTitle       string  `json:"full_title"`
	Pubisher        string  `json:"publisher"`
	PublicationDate string  `json:"publication_date"`
	Isbn            string  `json:"isbn"`
	Description     string  `json:"description"`
	Price           float64 `json:"price"`
	StockQuantity   int32   `json:"stock_quantity"`
	FrontCoverImage string  `json:"front_cover_image"`
	BackCoverImage  string  `json:"back_cover_image"`
}

func (server *Server) updateBook(ctx *gin.Context) {
	var req updateBookRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	convertedDate, err := util.ConvertStringToDate(req.PublicationDate)
	if err != nil {
		ctx.JSON(err.Code, ErrorResponse(err.Error))
		return
	}

	err = server.client.UpdateBook(&pb.UpdateBookRequest{
		Id:              req.Id,
		Title:           req.Title,
		FullTitle:       req.FullTitle,
		Publisher:       req.Pubisher,
		PublicationDate: convertedDate,
		Isbn:            req.Isbn,
		Description:     req.Description,
		Price:           req.Price,
		StockQuantity:   req.StockQuantity,
		FrontCoverImage: req.FrontCoverImage,
		BackCoverImage:  req.BackCoverImage,
	})

	if err != nil {
		ctx.JSON(err.Code, ErrorResponse(err.Error))
		return
	}

	ctx.JSON(http.StatusOK, struct {
		Message string `json:"message"`
	}{Message: "OK"})
}

type getAuthorRequest struct {
	Id string `uri:"id" binding:"required"`
}

func (server *Server) getAuthor(ctx *gin.Context) {
	var req getAuthorRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	author, err := server.client.GetAuthor(&pb.GetAuthorRequest{
		Id: req.Id,
	})

	if err != nil {
		ctx.JSON(err.Code, ErrorResponse(err.Error))
		return
	}

	ctx.JSON(http.StatusOK, author)
}

type getPublisherRequest struct {
	Id string `uri:"id" binding:"required"`
}

func (server *Server) getPublisher(ctx *gin.Context) {
	var req getPublisherRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	publisher, err := server.client.GetPublisher(&pb.GetPublisherRequest{
		Id: req.Id,
	})

	if err != nil {
		ctx.JSON(err.Code, ErrorResponse(err.Error))
		return
	}

	ctx.JSON(http.StatusOK, publisher)
}

type addAuthorRequest struct {
	FullName  string `json:"full_name" binding:"required"`
	BirthDate string `json:"birth_date" binding:"required"`
}

func (server *Server) addAuthor(ctx *gin.Context) {
	var req addAuthorRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	birthDate, err := util.ConvertStringToDate(req.BirthDate)
	if err != nil {
		ctx.JSON(err.Code, ErrorResponse(err.Error))
		return
	}

	author, err := server.client.AddNewAuthor(&pb.AddNewAuthorRequest{
		FullName:  req.FullName,
		BirthDate: birthDate,
	})

	if err != nil {
		ctx.JSON(err.Code, ErrorResponse(err.Error))
		return
	}

	ctx.JSON(http.StatusOK, author)
}

type addPublisherRequest struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
}

func (server *Server) addPublisher(ctx *gin.Context) {
	var req addPublisherRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	publisher, err := server.client.AddNewPublisher(&pb.AddNewPublisherRequest{
		Name:    req.Name,
		Address: req.Address,
	})

	if err != nil {
		ctx.JSON(err.Code, ErrorResponse(err.Error))
		return
	}

	ctx.JSON(http.StatusOK, publisher)
}
