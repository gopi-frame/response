package response

import (
	"encoding/json"
	"net/http"
)

// JSONResponse provides a convenient way to send JSON-encoded data
// as the response body in an HTTP request.
type JSONResponse struct {
	*Response
	data any
}

// SetContent sets response content
func (jsonResponse *JSONResponse) SetContent(data any) {
	jsonResponse.data = data
}

// ServeHTTP implements the http.Handler interface and writes the
// JSON-encoded response data to the ResponseWriter.
func (jsonResponse *JSONResponse) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	jsonBytes, err := json.Marshal(jsonResponse.data)
	if err != nil {
		panic(err)
	}
	jsonResponse.content = jsonBytes
	jsonResponse.SetHeader("content-type", "application/json")
	jsonResponse.Response.ServeHTTP(w, r)
}
