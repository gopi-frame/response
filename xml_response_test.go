package response

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http/httptest"
	"testing"
)

func TestXMLResponseSetContent(t *testing.T) {
	data := struct {
		Name string
		Age  int
	}{
		Name: "John",
		Age:  30,
	}

	xmlResponse := &XMLResponse{Response: New(200)}
	xmlResponse.SetContent(data)

	assert.Equal(t, data, xmlResponse.data)
}

func TestXMLResponseServeHTTP(t *testing.T) {
	data := struct {
		XMLName xml.Name `xml:"struct"`
		Name    string
		Age     int
	}{
		Name: "John",
		Age:  30,
	}

	xmlResponse := &XMLResponse{Response: New(200)}
	xmlResponse.SetContent(data)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/", nil)
	xmlResponse.ServeHTTP(recorder, request)

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
	assert.Contains(t, result.Header.Get("content-type"), "application/xml")
	assert.Equal(t, `<struct><Name>John</Name><Age>30</Age></struct>`, string(content))
}

func TestXMLResponseServeHTTPWithError(t *testing.T) {
	data := make(chan int)

	xmlResponse := &XMLResponse{Response: New(200)}
	xmlResponse.SetContent(data)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/", nil)

	assert.Panics(t, func() {
		xmlResponse.ServeHTTP(recorder, request)
	})
}
