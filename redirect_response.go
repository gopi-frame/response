package response

import (
	"fmt"
	"net/http"

	"github.com/gopi-frame/exception"
)

// RedirectResponse is a struct that facilitates sending an HTTP redirect response.
// It allows you to set the redirect location using the SetLocation method.
// When serving the response, it checks if the provided HTTP status code is within the valid range for redirection (between 300 and 308).
// If the status code is valid, it sends an HTTP redirect response to the specified location using the provided status code.
// If the status code is invalid for redirection, it panics with an appropriate exception message.
type RedirectResponse struct {
	*Response
	location string
}

// SetLocation sets the redirect location
func (redirectResponse *RedirectResponse) SetLocation(location string) *RedirectResponse {
	redirectResponse.location = location
	return redirectResponse
}

// ServeHTTP sends the response
func (redirectResponse *RedirectResponse) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if redirectResponse.statusCode < http.StatusMultipleChoices || redirectResponse.statusCode > http.StatusPermanentRedirect {
		panic(exception.New(fmt.Sprintf("can not redirect with HTTP status code `%d`", redirectResponse.statusCode)))
	}
	http.Redirect(w, r, redirectResponse.location, redirectResponse.statusCode)
}
