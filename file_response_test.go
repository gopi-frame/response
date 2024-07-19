package response

import (
	"io"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileResponse(t *testing.T) {
	f, err := os.CreateTemp(os.TempDir(), "test-file-response")
	if err != nil {
		assert.FailNow(t, err.Error())
	}
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
		if err := os.Remove(f.Name()); err != nil {
			panic(err)
		}
	}()
	data := []byte("Hello, World!")
	_, err = f.Write(data)
	if err != nil {
		assert.FailNow(t, err.Error())
	}
	response := New(200).File(f.Name())
	response.SetContentType("text/plain")
	response.SetContent(data)

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

func TestFileResponse_FileNotExists(t *testing.T) {
	response := New(200).File("not-exists.txt")
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/", nil)
	assert.Panics(t, func() { response.ServeHTTP(recorder, request) })
}
