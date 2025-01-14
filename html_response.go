package response

import (
	"bytes"
	"html/template"
	"net/http"
	"os"
)

type HtmlResponse struct {
	*Response
	html  string
	model map[string]any
}

func (h *HtmlResponse) SetHTML(html string) {
	h.html = html
}

func (h *HtmlResponse) LoadHtml(file string) error {
	content, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	h.html = string(content)
	return nil
}

func (h *HtmlResponse) SetModel(model map[string]any) {
	h.model = model
}

func (h *HtmlResponse) Assign(key string, value any) {
	h.model[key] = value
}

func (h *HtmlResponse) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	buf := new(bytes.Buffer)
	tpl := template.Must(template.New("html").Parse(h.html))
	if err := tpl.Execute(buf, h.model); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.Response.SetContent(buf.String())
	h.Response.ServeHTTP(w, r)
}
