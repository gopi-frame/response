package response

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http/httptest"
	"testing"
)

func TestJSONResponseSetContent(t *testing.T) {
	data := struct {
		Message string
		Value   int
	}{
		Message: "Hello, World!",
		Value:   42,
	}

	response := New(200).JSON()
	response.SetContent(data)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/", nil)
	response.ServeHTTP(recorder, request)

	result := recorder.Result()
	body, err := io.ReadAll(result.Body)
	assert.Nil(t, err)
	defer result.Body.Close()

	assert.Equal(t, 200, result.StatusCode)
	assert.Equal(t, "application/json", result.Header.Get("Content-Type"))
	assert.JSONEq(t, `{"Message":"Hello, World!","Value":42}`, string(body))
}

func TestJSONResponseMarshalError(t *testing.T) {
	data := make(chan int)

	response := New(200).JSON()
	response.SetContent(data)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/", nil)

	assert.Panics(t, func() {
		response.ServeHTTP(recorder, request)
	})
}

func TestJSONResponseNilData(t *testing.T) {
	response := New(200).JSON()
	response.SetContent(nil)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/", nil)
	response.ServeHTTP(recorder, request)

	result := recorder.Result()
	body, err := io.ReadAll(result.Body)
	assert.Nil(t, err)
	defer result.Body.Close()

	assert.Equal(t, 200, result.StatusCode)
	assert.Equal(t, "application/json", result.Header.Get("Content-Type"))
	assert.Equal(t, "null", string(body))
}
