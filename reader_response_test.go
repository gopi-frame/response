package response

import (
	"bytes"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReaderResponse(t *testing.T) {
	data := []byte("Hello, World!")
	reader := bytes.NewReader(data)

	response := New(200).Reader(reader)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/", nil)
	response.ServeHTTP(recorder, request)

	result := recorder.Result()
	body, err := io.ReadAll(result.Body)
	assert.Nil(t, err)
	defer result.Body.Close()

	assert.Equal(t, 200, result.StatusCode)
	assert.Equal(t, "text/plain; charset=utf-8", result.Header.Get("Content-Type"))
	assert.Equal(t, data, body)
}

func TestReaderResponseWithContentType(t *testing.T) {
	data := []byte("Hello, World!")
	reader := bytes.NewReader(data)

	response := New(200).Reader(reader)
	response.SetContentType("text/plain")
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/", nil)
	response.ServeHTTP(recorder, request)

	result := recorder.Result()
	body, err := io.ReadAll(result.Body)
	assert.Nil(t, err)
	defer result.Body.Close()

	assert.Equal(t, 200, result.StatusCode)
	assert.Equal(t, "text/plain", result.Header.Get("Content-Type"))
	assert.Equal(t, data, body)
}

func TestReaderResponseWithNilReader(t *testing.T) {
	response := New(200).Reader(nil)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/", nil)
	response.ServeHTTP(recorder, request)

	result := recorder.Result()
	body, err := io.ReadAll(result.Body)
	assert.Nil(t, err)
	defer result.Body.Close()

	assert.Equal(t, 200, result.StatusCode)
	assert.Equal(t, "application/octet-stream", result.Header.Get("Content-Type"))
	assert.Empty(t, body)
}
