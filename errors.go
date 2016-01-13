package restclient

import (
	"github.com/juju/errgo"
)

var (
	ForbiddenError       = errgo.New("forbidden")
	NotFoundError        = errgo.New("not found")
	InternalServerError  = errgo.New("internal server error")
	InvalidArgumentError = errgo.New("invalid argument")
	UnauthorizedError    = errgo.New("unauthorized")
	maskAny              = errgo.MaskFunc(errgo.Any)
)

func IsInvalidArgument(err error) bool {
	return errgo.Cause(err) == InvalidArgumentError
}

func IsNotFound(err error) bool {
	return errgo.Cause(err) == NotFoundError
}

func IsInternalServer(err error) bool {
	return errgo.Cause(err) == InternalServerError
}

type ErrorResponse struct {
	TheError struct {
		Message string `json:"message"`
	} `json:"error"`
}

func (er *ErrorResponse) Error() string {
	return er.TheError.Message
}

func NewErrorResponse(message string) ErrorResponse {
	er := ErrorResponse{}
	er.TheError.Message = message
	return er
}
