package util

import (
	"errors"
	"time"

	httpError "github.com/khiemta03/bookstore-be/api-getway/internal/error"
	"google.golang.org/genproto/googleapis/type/date"
)

func ConvertStringToDate(dateStr string) (*date.Date, *httpError.HTTPError) {
	convertedDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil, &httpError.HTTPError{
			Code:  400,
			Error: errors.New("date format is invalid"),
		}
	}

	return &date.Date{
		Year:  int32(convertedDate.Year()),
		Month: int32(convertedDate.Month()),
		Day:   int32(convertedDate.Day()),
	}, nil
}
