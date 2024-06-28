package httpError

import (
	"errors"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type HTTPError struct {
	Code  int   `json:"code"`
	Error error `json:"error"`
}

func MapGRPCErrorToHTTPError(err error) *HTTPError {
	var httpErr HTTPError
	if st, ok := status.FromError(err); ok {
		httpErr.Error = errors.New(st.Message())
		switch st.Code() {
		case codes.Internal:
			httpErr.Code = http.StatusInternalServerError
		case codes.AlreadyExists:
			httpErr.Code = http.StatusBadRequest
		case codes.InvalidArgument:
			httpErr.Code = http.StatusBadRequest
		case codes.NotFound:
			httpErr.Code = http.StatusNotFound
		default:
			httpErr.Code = http.StatusBadRequest
		}
	}

	return &httpErr
}
