package api

import (
	"context"

	db "github.com/khiemta03/bookstore-be/book-service/internal/database/sqlc"
	ce "github.com/khiemta03/bookstore-be/book-service/internal/error"
	pb "github.com/khiemta03/bookstore-be/book-service/internal/grpc/gen/book"
	utils "github.com/khiemta03/bookstore-be/book-service/internal/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) AddNewBook(ctx context.Context, req *pb.AddNewBookRequest) (*pb.Book, error) {
	convertedPublisherId, err := utils.ConvertToUUID(req.GetPublisher())
	if err != nil {
		// TODO: send err to log service
		return nil, status.Errorf(codes.InvalidArgument, ce.ErrInvalidAgrumentStr)
	}

	convertedAuthorIds, err := utils.ConvertListToUUIDs(req.Authors)
	if err != nil {
		// TODO: send err to log service
		return nil, status.Errorf(codes.InvalidArgument, ce.ErrInvalidAgrumentStr)
	}

	txResult, cerr := server.store.AddNewBookTx(ctx, db.AddNewBookTxParams{
		Title:           req.GetTitle(),
		FullTitle:       req.GetFullTitle(),
		Publisher:       *convertedPublisherId,
		PublicationDate: convertDateToTime(req.GetPublicationDate()),
		Isbn:            req.GetIsbn(),
		Description:     convertStringToNullString(req.GetDescription()),
		Price:           req.GetPrice(),
		StockQuantity:   req.GetStockQuantity(),
		FrontCoverImage: convertStringToNullString(req.GetFrontCoverImage()),
		BackCoverImage:  convertStringToNullString(req.GetBackCoverImage()),
		Authors:         convertedAuthorIds,
	})
	if !cerr.IsNil() {
		// TODO: send err to log service
		switch cerr.CustomErr {
		case ce.ErrInternalServer:
			return nil, status.Errorf(codes.NotFound, ce.ErrInternalServerStr)
		default:
			return nil, status.Errorf(codes.InvalidArgument, cerr.CustomErr.Error())
		}
	}

	res := convertBook(txResult.Book)
	res.Publisher = convertPublisher(txResult.Publisher)
	for _, author := range txResult.Authors {
		res.Authors = append(res.Authors, convertAuthor(author))
	}

	return res, nil
}
