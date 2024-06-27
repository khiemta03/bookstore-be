package api

import (
	"database/sql"
	"time"

	db "github.com/khiemta03/bookstore-be/book-service/internal/database/sqlc"
	pb "github.com/khiemta03/bookstore-be/book-service/internal/grpc/gen/book"
	"google.golang.org/genproto/googleapis/type/date"
)

func convertBook(user db.BOOK) *pb.Book {
	return &pb.Book{
		Id:              user.ID.String(),
		Title:           user.Title,
		FullTitle:       user.FullTitle,
		PublicationDate: convertTimeToDate(user.PublicationDate),
		Isbn:            user.Isbn,
		Description:     user.Description.String,
		Price:           user.Price,
		StockQuantity:   user.StockQuantity,
		FrontCoverImage: user.FrontCoverImage.String,
		BackCoverImage:  user.BackCoverImage.String,
	}
}

func convertAuthor(author db.AUTHOR) *pb.Author {
	return &pb.Author{
		Id:        author.ID.String(),
		FullName:  author.FullName,
		BirthDate: convertNullTimeToDate(author.Birthdate),
	}
}

func convertPublisher(publisher db.PUBLISHER) *pb.Publisher {
	return &pb.Publisher{
		Id:      publisher.ID.String(),
		Name:    publisher.Name,
		Address: publisher.Address.String,
	}
}

func convertDateToTime(date *date.Date) time.Time {
	if date == nil {
		return time.Time{}
	}

	return time.Date(
		int(date.Year),
		time.Month(date.Month),
		int(date.Day),
		0, 0, 0, 0,
		time.UTC,
	)
}

func convertDateToNullTime(date *date.Date) sql.NullTime {
	if date == nil {
		return sql.NullTime{}
	}

	return sql.NullTime{
		Time:  convertDateToTime(date),
		Valid: true,
	}
}

func convertTimeToDate(time time.Time) *date.Date {
	return &date.Date{
		Year:  int32(time.Year()),
		Month: int32(time.Month()),
		Day:   int32(time.Day()),
	}
}

func convertNullTimeToDate(ntm sql.NullTime) *date.Date {
	if ntm.Valid {
		return convertTimeToDate(ntm.Time)
	}

	return nil
}

func convertStringToNullString(str string) sql.NullString {
	if str == "" {
		return sql.NullString{
			String: "",
			Valid:  false,
		}
	}
	return sql.NullString{
		String: str,
		Valid:  true,
	}
}
