// Copyright (c) 2016 Epracom Advies.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package restkit

import (
	"net/http"

	"github.com/juju/errgo"
)

var (
	ForbiddenError          = newErrorResponseWithStatusCodeFunc(http.StatusForbidden)
	InternalServerError     = newErrorResponseWithStatusCodeFunc(http.StatusInternalServerError)
	BadRequestError         = newErrorResponseWithStatusCodeFunc(http.StatusBadRequest)
	NotFoundError           = newErrorResponseWithStatusCodeFunc(http.StatusNotFound)
	PreconditionFailedError = newErrorResponseWithStatusCodeFunc(http.StatusPreconditionFailed)
	UnauthorizedError       = newErrorResponseWithStatusCodeFunc(http.StatusUnauthorized)
	maskAny                 = errgo.MaskFunc(errgo.Any)
)

func IsStatusBadRequest(err error) bool {
	return isErrorResponseWithStatusCode(err, http.StatusBadRequest)
}

func IsStatusForbidden(err error) bool {
	return isErrorResponseWithStatusCode(err, http.StatusForbidden)
}

func IsStatusInternalServer(err error) bool {
	return isErrorResponseWithStatusCode(err, http.StatusInternalServerError)
}

func IsStatusNotFound(err error) bool {
	return isErrorResponseWithStatusCode(err, http.StatusNotFound)
}

func IsStatusPreconditionFailed(err error) bool {
	return isErrorResponseWithStatusCode(err, http.StatusPreconditionFailed)
}

func IsStatusUnauthorizedError(err error) bool {
	return isErrorResponseWithStatusCode(err, http.StatusUnauthorized)
}

type ErrorResponse struct {
	TheError struct {
		Message string `json:"message,omitempty"`
		Code    int    `json:"code,omitempty"`
	} `json:"error"`

	// HTTP status code
	statusCode int `json:"-"`
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

func IsErrorResponseWithCodeFunc(code int) func(error) bool {
	return func(err error) bool {
		return IsErrorResponseWithCode(err, code)
	}
}

func isErrorResponseWithStatusCode(err error, statusCode int) bool {
	if er, ok := errgo.Cause(err).(*ErrorResponse); ok {
		return er.statusCode == statusCode
	}
	return false
}

func NewErrorResponse(message string, code int) error {
	er := &ErrorResponse{}
	er.TheError.Message = message
	er.TheError.Code = code
	return er
}

func newErrorResponseWithStatusCodeFunc(statusCode int) func(string, int) error {
	return func(message string, code int) error {
		er := &ErrorResponse{}
		er.TheError.Message = message
		er.TheError.Code = code
		er.statusCode = statusCode
		return er
	}
}
