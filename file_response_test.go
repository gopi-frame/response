package response

import (
	"io"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileResponse(t *testing.T) {
	tempDir := os.TempDir()
	tempFile, err := os.CreateTemp(tempDir, "test_response")
	assert.Nil(t, err)
	if _, err := tempFile.Write([]byte("helloworld")); err != nil {
		assert.FailNow(t, err.Error())
	}
	response := New(200).File(tempFile.Name())
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/", nil)
	response.ServeHTTP(recorder, request)
	result := recorder.Result()
	body := result.Body
	content, err := io.ReadAll(body)
	assert.Nil(t, err)
	defer func() {
		if err := body.Close(); err != nil {
			panic(err)
		}
	}()
	assert.Equal(t, 200, result.StatusCode)
	assert.Contains(t, result.Header.Get("content-type"), "text/plain")
	assert.Equal(t, "helloworld", string(content))
}
