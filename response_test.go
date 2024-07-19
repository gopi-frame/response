package response

import (
	"bytes"
	"encoding/xml"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"net/http/httptest"
)

func TestNewResponse(t *testing.T) {
	response := New(200)
	assert.Equal(t, 200, response.StatusCode())
	assert.Nil(t, response.Content())
	assert.Equal(t, "", response.Header("Content-Type"))
}

func TestResponseSetStatusCode(t *testing.T) {
	response := New(200)
	response.SetStatusCode(404)
	assert.Equal(t, 404, response.StatusCode())
}

func TestResponseSetHeader(t *testing.T) {
	response := New(200)
	response.SetHeader("X-Custom-Header", "custom-value")
	assert.Equal(t, "custom-value", response.Header("X-Custom-Header"))
}

func TestResponseServeHttp(t *testing.T) {
	response := New(200, "Hello, World!")
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/", nil)
	response.ServeHTTP(recorder, request)

	result := recorder.Result()
	assert.Equal(t, 200, result.StatusCode)
	content, err := io.ReadAll(result.Body)
	if err != nil {
		assert.FailNow(t, err.Error())
	}
	defer result.Body.Close()
	assert.Equal(t, "Hello, World!", string(content))
}

func TestResponse_JSON(t *testing.T) {
	response := New(200).JSON(map[string]interface{}{"message": "Hello, World!"})
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/", nil)
	response.ServeHTTP(recorder, request)
	assert.Equal(t, "application/json", response.Header("Content-Type"))
	result := recorder.Result()
	assert.Equal(t, 200, result.StatusCode)
	assert.Equal(t, "application/json", result.Header.Get("Content-Type"))
	content, err := io.ReadAll(result.Body)
	if err != nil {
		assert.FailNow(t, err.Error())
	}
	defer result.Body.Close()
	assert.JSONEq(t, `{"message":"Hello, World!"}`, string(content))
}

func TestResponse_XML(t *testing.T) {
	type Message struct {
		XMLName xml.Name `xml:"response"`
		Message string   `xml:"message"`
	}
	response := New(200).XML(Message{Message: "Hello, World!"})
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/", nil)
	response.ServeHTTP(recorder, request)
	assert.Equal(t, "application/xml", response.Header("Content-Type"))
	result := recorder.Result()
	assert.Equal(t, 200, result.StatusCode)
	assert.Equal(t, "application/xml", result.Header.Get("Content-Type"))
	content, err := io.ReadAll(result.Body)
	if err != nil {
		assert.FailNow(t, err.Error())
	}
	defer result.Body.Close()
	assert.Equal(t, `<response><message>Hello, World!</message></response>`, string(content))
}

func TestResponse_Redirect(t *testing.T) {
	response := New(302).Redirect("https://example.com")
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/", nil)
	response.ServeHTTP(recorder, request)
	result := recorder.Result()
	assert.Equal(t, 302, result.StatusCode)
	assert.Equal(t, "https://example.com", result.Header.Get("Location"))
}

func TestResponse_Reader(t *testing.T) {
	data := []byte("Hello, World!")
	reader := bytes.NewReader(data)
	response := New(200).Reader(reader)
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/", nil)
	response.ServeHTTP(recorder, request)
	result := recorder.Result()
	assert.Equal(t, 200, result.StatusCode)
	assert.Equal(t, "text/plain; charset=utf-8", result.Header.Get("Content-Type"))
	content, err := io.ReadAll(result.Body)
	if err != nil {
		assert.FailNow(t, err.Error())
	}
	defer result.Body.Close()
	assert.Equal(t, data, content)
}

func TestResponse_File(t *testing.T) {
	f, err := os.CreateTemp(os.TempDir(), "response_test")
	if err != nil {
		assert.FailNow(t, err.Error())
	}
	defer func() {
		if err := f.Close(); err != nil {
			assert.FailNow(t, err.Error())
		}
		if err := os.Remove(f.Name()); err != nil {
			assert.FailNow(t, err.Error())
		}
	}()
	_, err = f.Write([]byte("Hello, World!\n"))
	if err != nil {
		assert.FailNow(t, err.Error())
	}
	response := New(200).File(f.Name())
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/", nil)
	response.ServeHTTP(recorder, request)
	result := recorder.Result()
	assert.Equal(t, 200, result.StatusCode)
	assert.Equal(t, "text/plain; charset=utf-8", result.Header.Get("Content-Type"))
	content, err := io.ReadAll(result.Body)
	if err != nil {
		assert.FailNow(t, err.Error())
	}
	defer result.Body.Close()
	assert.Equal(t, "Hello, World!\n", string(content))
}

func TestResponse_Stream(t *testing.T) {
	i := 0
	response := New(200).Stream(func(w io.Writer) bool {
		_, err := w.Write([]byte("Hello, World!\n"))
		if err != nil {
			assert.FailNow(t, err.Error())
		}
		i++
		return i < 3
	})
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/", nil)
	response.ServeHTTP(recorder, request)
	result := recorder.Result()
	assert.Equal(t, 200, result.StatusCode)
	assert.Equal(t, "text/plain; charset=utf-8", result.Header.Get("Content-Type"))
	content, err := io.ReadAll(result.Body)
	if err != nil {
		assert.FailNow(t, err.Error())
	}
	defer result.Body.Close()
	assert.Equal(t, "Hello, World!\nHello, World!\nHello, World!\n", string(content))
}
