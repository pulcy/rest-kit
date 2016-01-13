package restclient

import (
	"encoding/json"
	"net/http"
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
