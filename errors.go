package restkit

import (
	"github.com/juju/errgo"
)

var (
	ForbiddenError       = errgo.New("forbidden")
	InternalServerError  = errgo.New("internal server error")
	InvalidArgumentError = errgo.New("invalid argument")
	NotFoundError        = errgo.New("not found")
	UnauthorizedError    = errgo.New("unauthorized")
	maskAny              = errgo.MaskFunc(errgo.Any)
)

func IsForbidden(err error) bool {
	return errgo.Cause(err) == ForbiddenError
}

func IsInternalServer(err error) bool {
	return errgo.Cause(err) == InternalServerError
}

func IsInvalidArgument(err error) bool {
	return errgo.Cause(err) == InvalidArgumentError
}

func IsNotFound(err error) bool {
	return errgo.Cause(err) == NotFoundError
}

func IsUnauthorizedError(err error) bool {
	return errgo.Cause(err) == UnauthorizedError
}

type ErrorResponse struct {
	TheError struct {
		Message string `json:"message,omitempty"`
		Code    int    `json:"code,omitempty"`
	} `json:"error"`
}

func (er *ErrorResponse) Error() string {
	return er.TheError.Message
}

func IsErrorResponseWithCode(err error, code int) bool {
	if er, ok := errgo.Cause(err).(*ErrorResponse); ok {
		return er.TheError.Code == code
	}
	return false
}

func NewErrorResponse(message string, code int) error {
	er := ErrorResponse{}
	er.TheError.Message = message
	er.TheError.Code = code
	return &er
}
