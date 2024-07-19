package response

import (
	"io"
	"net/http"

	"github.com/gabriel-vasile/mimetype"
)

// ReaderResponse is a struct that allows you to send the contents of an [io.Reader] as the response body in an HTTP request.
// It provides methods to set the reader itself (SetReader) and the content type of the response (`SetContentType`).
// If the content type is not set explicitly, it attempts to detect the MIME type based on the reader's contents.
// This struct is useful when you need to stream data from a reader, such as a file or an in-memory buffer, as the response body.
type ReaderResponse struct {
	*Response
	contentType string
	reader      io.Reader
}

// SetReader sets the reader
func (readerResponse *ReaderResponse) SetReader(reader io.Reader) *ReaderResponse {
	readerResponse.reader = reader
	return readerResponse
}

// SetContentType sets response Content-Type header
func (readerResponse *ReaderResponse) SetContentType(contentType string) *ReaderResponse {
	readerResponse.contentType = contentType
	return readerResponse
}

// ServeHTTP sends the response
func (readerResponse *ReaderResponse) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// set cookies
	for _, cookie := range readerResponse.cookies {
		http.SetCookie(w, cookie)
	}
	// set headers
	for key, value := range readerResponse.headers {
		w.Header()[key] = value
	}
	// set content type
	if readerResponse.contentType != "" {
		w.Header().Set("content-type", readerResponse.contentType)
	} else if readerResponse.reader != nil {
		mime, _ := mimetype.DetectReader(readerResponse.reader)
		w.Header().Set("content-type", mime.String())
		// rewind reader
		if _, err := readerResponse.reader.(io.ReadSeeker).Seek(0, 0); err != nil {
			panic(err)
		}
	} else {
		w.Header().Set("content-type", "application/octet-stream")
	}
	// set http status code
	w.WriteHeader(readerResponse.statusCode)
	if readerResponse.reader == nil {
		return
	}
	if _, err := io.Copy(w, readerResponse.reader); err != nil {
		panic(err)
	}
}
