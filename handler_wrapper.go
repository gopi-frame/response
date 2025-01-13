package response

import (
	"net/http"

	"github.com/gopi-frame/exception"
)

type HandlerWrapper struct {
	handler http.Handler
}

func NewHandlerWrapper(handler http.Handler) *HandlerWrapper {
	return &HandlerWrapper{handler}
}

func (hw *HandlerWrapper) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	hw.handler.ServeHTTP(writer, request)
}

func (hw *HandlerWrapper) SetStatusCode(_ int) {
	panic(exception.NewUnsupportedException("not supported yet"))
}

func (hw *HandlerWrapper) StatusCode() int {
	panic(exception.NewUnsupportedException("not supported yet"))
}

func (hw *HandlerWrapper) SetContent(_ any) {
	panic(exception.NewUnsupportedException("not supported yet"))
}

func (hw *HandlerWrapper) Content() any {
	panic(exception.NewUnsupportedException("not supported yet"))
}

func (hw *HandlerWrapper) SetHeader(_, _ string, _ ...bool) {
	panic(exception.NewUnsupportedException("not supported yet"))
}

func (hw *HandlerWrapper) SetHeaders(_ map[string]string) {
	panic(exception.NewUnsupportedException("not supported yet"))
}

func (hw *HandlerWrapper) HasHeader(_ string) bool {
	panic(exception.NewUnsupportedException("not supported yet"))
}

func (hw *HandlerWrapper) Header(_ string) string {
	panic(exception.NewUnsupportedException("not supported yet"))
}

func (hw *HandlerWrapper) Headers() http.Header {
	panic(exception.NewUnsupportedException("not supported yet"))
}

func (hw *HandlerWrapper) SetCookie(_ *http.Cookie) {
	panic(exception.NewUnsupportedException("not supported yet"))
}

func (hw *HandlerWrapper) Cookies() []*http.Cookie {
	panic(exception.NewUnsupportedException("not supported yet"))
}
