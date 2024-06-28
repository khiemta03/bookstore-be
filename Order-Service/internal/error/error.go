// custom error package
package ce

import "errors"

var (
	ErrInvalidAgrumentStr          = "invalid agrument error"
	ErrItemExistedStr              = "item has been added into cart"
	ErrOrderNotFoundStr            = "order not found error"
	ErrOrderDetailNotFoundStr      = "order detail not found error"
	ErrDiscountNotFoundStr         = "discount not found error"
	ErrShoppingCartItemNotFoundStr = "shopping cart item not found error"
	ErrNonGRPCStr                  = "non-gRPC error"
	ErrInternalServerStr           = "oops something went wrong"
	ErrUnknownGRPCStr              = "unknown gRPC error"
)

var (
	ErrInvalidAgrument          = errors.New(ErrInvalidAgrumentStr)
	ErrItemExisted              = errors.New(ErrItemExistedStr)
	ErrOrderNotFound            = errors.New(ErrOrderNotFoundStr)
	ErrDiscountNotFound         = errors.New(ErrDiscountNotFoundStr)
	ErrOrderDetailNotFound      = errors.New(ErrOrderDetailNotFoundStr)
	ErrShoppingCartItemNotFound = errors.New(ErrShoppingCartItemNotFoundStr)
	ErrNonGRPC                  = errors.New(ErrNonGRPCStr)
	ErrInternalServer           = errors.New(ErrInternalServerStr)
	ErrUnknownGRPC              = errors.New(ErrUnknownGRPCStr)
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

func OrderNotFoundError(err error) CustomError {
	return CustomError{
		OriginalErr: err,
		CustomErr:   ErrOrderNotFound,
	}
}

func DiscountNotFoundError(err error) CustomError {
	return CustomError{
		OriginalErr: err,
		CustomErr:   ErrDiscountNotFound,
	}
}

func OrderDetailNotFoundError(err error) CustomError {
	return CustomError{
		OriginalErr: err,
		CustomErr:   ErrOrderDetailNotFound,
	}
}

func ShoppingCartItemNotFoundError(err error) CustomError {
	return CustomError{
		OriginalErr: err,
		CustomErr:   ErrShoppingCartItemNotFound,
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
