package response

import (
	"net/http"
	"os"
)

// FileResponse is used to send a file response
type FileResponse struct {
	*ReaderResponse
	filename string
}

// SetFile sets the filename
func (fileResponse *FileResponse) SetFile(filename string) *FileResponse {
	fileResponse.filename = filename
	return fileResponse
}

// ServeHTTP reads the file content and sends it
func (fileResponse *FileResponse) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open(fileResponse.filename)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()
	fileResponse.SetReader(f)
	fileResponse.ReaderResponse.ServeHTTP(w, r)
}
