package response

import (
	"context"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http/httptest"
	"testing"
)

func TestStreamedResponseSetStep(t *testing.T) {
	response := New(200).Stream(nil)
	stepCalled := false
	step := func(w io.Writer) bool {
		stepCalled = true
		return false
	}
	response.SetStep(step)
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/", nil)
	response.ServeHTTP(recorder, request)
	assert.True(t, stepCalled)
}

func TestStreamedResponseServeHTTPWithoutStep(t *testing.T) {
	response := New(200).Stream(nil)
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/", nil)
	response.ServeHTTP(recorder, request)
	result := recorder.Result()
	assert.Equal(t, 200, result.StatusCode)
}

func TestStreamedResponseServeHTTPWithCanceledContext(t *testing.T) {
	response := New(200).Stream(nil)
	stepCalled := false
	step := func(w io.Writer) bool {
		stepCalled = true
		return true
	}
	response.SetStep(step)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	request := httptest.NewRequest("GET", "/", nil)
	request = request.WithContext(ctx)
	recorder := httptest.NewRecorder()
	response.ServeHTTP(recorder, request)
	assert.False(t, stepCalled)
}

func TestStreamedResponseServeHTTPWithMultipleSteps(t *testing.T) {
	response := New(200).Stream(nil)
	stepCount := 0
	step := func(w io.Writer) bool {
		stepCount++
		_, err := w.Write([]byte("Hello"))
		if err != nil {
			return false
		}
		if stepCount < 3 {
			return true
		}
		return false
	}
	response.SetStep(step)
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/", nil)
	response.ServeHTTP(recorder, request)
	assert.Equal(t, 3, stepCount)
	assert.Equal(t, 200, recorder.Result().StatusCode)
	content, err := io.ReadAll(recorder.Result().Body)
	if err != nil {
		assert.FailNow(t, err.Error())
	}
	defer recorder.Result().Body.Close()
	assert.Equal(t, "HelloHelloHello", string(content))
}
