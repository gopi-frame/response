package response

import (
	"encoding/xml"
	"net/http"
)

// XMLResponse used to send XML-encoded data
type XMLResponse struct {
	*Response
	data any
}

// SetContent sets response body content
func (xmlResponse *XMLResponse) SetContent(data any) {
	xmlResponse.data = data
}

// ServeHTTP sends the response
func (xmlResponse *XMLResponse) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	xmlBytes, err := xml.Marshal(xmlResponse.data)
	if err != nil {
		panic(err)
	}
	xmlResponse.content = xmlBytes
	xmlResponse.SetHeader("content-type", "application/xml")
	xmlResponse.Response.ServeHTTP(w, r)
}
