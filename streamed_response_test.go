package response

import (
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStreamedResponse(t *testing.T) {
	var i int
	response := New(200).Stream(func(w io.Writer) bool {
		if _, err := w.Write([]byte(fmt.Sprint(i))); err != nil {
			assert.FailNow(t, err.Error())
			return false
		}
		i++
		return i < 10
	})
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/", nil)
	response.ServeHTTP(recorder, request)
	result := recorder.Result()
	body := result.Body
	content, err := io.ReadAll(body)
	if err != nil {
		assert.FailNow(t, err.Error())
	}
	defer func() {
		if err := body.Close(); err != nil {
			assert.FailNow(t, err.Error())
		}
	}()
	assert.Equal(t, "0123456789", string(content))
}
