package response

import (
	"io"
	"net/http"
)

// StreamedResponse used to send a streamed response
type StreamedResponse struct {
	*Response
	step func(w io.Writer) bool
}

// SetStep sets the step func
func (streamed *StreamedResponse) SetStep(step func(w io.Writer) bool) *StreamedResponse {
	streamed.step = step
	return streamed
}

// ServeHTTP sends the response
func (streamed *StreamedResponse) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	select {
	case <-ctx.Done():
	default:
		for {
			if streamed.step == nil {
				break
			}
			if !streamed.step(w) {
				break
			}
		}
		w.(http.Flusher).Flush()
	}
}
