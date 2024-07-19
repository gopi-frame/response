package response

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRedirectResponse(t *testing.T) {
	t.Run("RedirectToURL", func(t *testing.T) {
		url := "https://example.com"
		response := New(302).Redirect(url)

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("GET", "/", nil)
		response.ServeHTTP(recorder, request)

		result := recorder.Result()
		assert.Equal(t, http.StatusFound, result.StatusCode)
		assert.Equal(t, url, result.Header.Get("Location"))
	})

	t.Run("RedirectToEmptyURL", func(t *testing.T) {
		response := New(302).Redirect("")

		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("GET", "/", nil)
		response.ServeHTTP(recorder, request)

		result := recorder.Result()
		assert.Equal(t, http.StatusFound, result.StatusCode)
		assert.Equal(t, "/", result.Header.Get("Location"))
	})

	t.Run("Redirect with invalid status code", func(t *testing.T) {
		url := "https://example.com"
		response := New(400).Redirect(url)
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest("GET", "/", nil)
		assert.Panics(t, func() {
			response.ServeHTTP(recorder, request)
		})
	})
}
