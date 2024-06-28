// custom error package
package ce

import "errors"

var (
	ErrInvalidAgrumentStr   = "invalid agrument error"
	ErrBookNotFoundStr      = "book not found error"
	ErrAuthorNotFoundStr    = "author not found error"
	ErrPublisherNotFoundStr = "publisher not found error"
	ErrNonGRPCStr           = "non-gRPC error"
	ErrAlreadyExistsStr     = "user already exists error"
	ErrOutOfStockStr        = "book out of stock"
	ErrInternalServerStr    = "internal server error"
	ErrUnknownGRPCStr       = "unknown gRPC error"
)

var (
	ErrInvalidAgrument   = errors.New(ErrInvalidAgrumentStr)
	ErrBookNotFound      = errors.New(ErrBookNotFoundStr)
	ErrAuthorNotFound    = errors.New(ErrAuthorNotFoundStr)
	ErrPublisherNotFound = errors.New(ErrPublisherNotFoundStr)
	ErrNonGRPC           = errors.New(ErrNonGRPCStr)
	ErrAlreadyExists     = errors.New(ErrAlreadyExistsStr)
	ErrOutOfStock        = errors.New(ErrOutOfStockStr)
	ErrInternalServer    = errors.New(ErrInternalServerStr)
	ErrUnknownGRPC       = errors.New(ErrUnknownGRPCStr)
)

type CustomError struct {
	OriginalErr error `json:"original_error"`
	CustomErr   error `json:"custom_error"`
}

func NilCustomError() CustomError {
	return CustomError{
		OriginalErr: nil,
		CustomErr:   nil,
	}
}

func InvalidAgrumentError(err error) CustomError {
	return CustomError{
		OriginalErr: err,
		CustomErr:   ErrInvalidAgrument,
	}
}

func BookNotFoundError(err error) CustomError {
	return CustomError{
		OriginalErr: err,
		CustomErr:   ErrBookNotFound,
	}
}

func OutOfStockError(err error) CustomError {
	return CustomError{
		OriginalErr: err,
		CustomErr:   ErrOutOfStock,
	}
}

func AuthorNotFoundError(err error) CustomError {
	return CustomError{
		OriginalErr: err,
		CustomErr:   ErrAuthorNotFound,
	}
}

func PublisherNotFoundError(err error) CustomError {
	return CustomError{
		OriginalErr: err,
		CustomErr:   ErrPublisherNotFound,
	}
}

func InternalServerError(err error) CustomError {
	return CustomError{
		OriginalErr: err,
		CustomErr:   ErrInternalServer,
	}
}

func (e *CustomError) IsNil() bool {
	return e.OriginalErr == nil
}
