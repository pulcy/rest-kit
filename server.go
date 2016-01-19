package restkit

import (
	"encoding/json"
	"net/http"

	"github.com/juju/errgo"
)

// JSON creates a application/json content-type header, sets the given HTTP
// status code and encodes the given result object to the response writer.
func JSON(resp http.ResponseWriter, result interface{}, code int) error {
	resp.Header().Add("Content-Type", "application/json")
	resp.WriteHeader(code)
	if result != nil {
		return maskAny(json.NewEncoder(resp).Encode(result))
	}
	return nil
}

// Text creates a text/plain content-type header, sets the given HTTP
// status code and writes the given content to the response writer.
func Text(resp http.ResponseWriter, content string, code int) error {
	resp.Header().Add("Content-Type", "text/plain")
	resp.WriteHeader(code)
	_, err := resp.Write([]byte(content))
	return maskAny(err)
}

// Html creates a text/html content-type header, sets the given HTTP
// status code and writes the given content to the response writer.
func Html(resp http.ResponseWriter, content string, code int) error {
	resp.Header().Add("Content-Type", "text/html")
	resp.WriteHeader(code)
	_, err := resp.Write([]byte(content))
	return maskAny(err)
}

// Error sends an error message back to the given response writer.
func Error(resp http.ResponseWriter, err error) error {
	code := http.StatusBadRequest
	var er *ErrorResponse

	if erX, ok := err.(*ErrorResponse); ok {
		er = erX
	} else if erX, ok := errgo.Cause(err).(*ErrorResponse); ok {
		er = erX
	} else {
		er = &ErrorResponse{}
		er.TheError.Message = err.Error()
		er.TheError.Code = -1
	}
	if er.statusCode != 0 {
		code = er.statusCode
	}
	return maskAny(JSON(resp, er, code))
}
