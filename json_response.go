package response

import (
	"encoding/json"
	"net/http"
)

// JSONResponse used to response json format data
type JSONResponse struct {
	*Response
	data any
}

// SetContent sets response body content
func (jsonResponse *JSONResponse) SetContent(data any) {
	jsonResponse.data = data
}

// ServeHTTP sends the response
func (jsonResponse *JSONResponse) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	jsonBytes, err := json.Marshal(jsonResponse.data)
	if err != nil {
		panic(err)
	}
	jsonResponse.content = jsonBytes
	jsonResponse.SetHeader("content-type", "application/json")
	jsonResponse.Response.ServeHTTP(w, r)
}
