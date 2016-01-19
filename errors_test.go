package restkit

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestError(t *testing.T) {
	tests := []struct {
		Error  error
		Tester func(error) bool
	}{
		{
			Error:  BadRequestError("foo", 0),
			Tester: IsStatusBadRequest,
		},
		{
			Error:  ForbiddenError("abc", 1),
			Tester: IsStatusForbidden,
		},
		{
			Error:  InternalServerError("", 0),
			Tester: IsStatusInternalServer,
		},
		{
			Error:  NotFoundError("id", 1),
			Tester: IsStatusNotFound,
		},
		{
			Error:  PreconditionFailedError("arg", 1),
			Tester: IsStatusPreconditionFailed,
		},
		{
			Error:  UnauthorizedError("user", 123456789),
			Tester: IsStatusUnauthorizedError,
		},
		{
			Error:  NewErrorResponse("test", 123),
			Tester: func(err error) bool { return IsErrorResponseWithCode(err, 123) },
		},
		{
			Error:  maskAny(BadRequestError("val", 0)),
			Tester: IsStatusBadRequest,
		},
		{
			Error:  maskAny(ForbiddenError("method", 1)),
			Tester: IsStatusForbidden,
		},
		{
			Error:  maskAny(InternalServerError("func", 9)),
			Tester: IsStatusInternalServer,
		},
		{
			Error:  maskAny(NotFoundError("key", 7)),
			Tester: IsStatusNotFound,
		},
		{
			Error:  maskAny(PreconditionFailedError("condition-x", 3)),
			Tester: IsStatusPreconditionFailed,
		},
		{
			Error:  maskAny(UnauthorizedError("group", 2)),
			Tester: IsStatusUnauthorizedError,
		},
		{
			Error:  maskAny(NewErrorResponse("test", 123)),
			Tester: func(err error) bool { return IsErrorResponseWithCode(err, 123) },
		},
	}

	for _, test := range tests {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			Error(w, test.Error)
		}))
		defer ts.Close()

		url, err := url.Parse(ts.URL)
		if err != nil {
			t.Fatalf("Failed to parse '%s': %#v", ts.URL, err)
		}
		rc := NewRestClient(url)
		err = rc.Request("GET", "/", nil, nil, nil)
		if !test.Tester(err) {
			t.Fatalf("Error test failed for %#v", test.Error)
		}
	}
}
